package payment

import (
	"context"
	"errors"
	"github.com/dtm-labs/dtmdriver"
	"github.com/zeromicro/go-zero/zrpc"
	"go-common/tool"
	"go-zero-base/utils/xerr"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"greet-pb/order/orderclient"
	"greet-pb/order/types/order"
	"greet-pb/payment/types/payment"
	"payment/cmd/business"
	"payment/cmd/dao/model"

	"payment/cmd/api/internal/svc"
	"payment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	// 下面这行导入gozero的dtm驱动
	"github.com/dtm-labs/client/dtmgrpc"
	_ "github.com/dtm-labs/driver-gozero"
)

type OrderPaymentNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderPaymentNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPaymentNoticeLogic {
	return &OrderPaymentNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderPaymentNoticeLogic) OrderPaymentNotice(req *types.PaymentNoticeReq) (resp *types.PaymentNoticeResp, err error) {
	logx.WithContext(l.ctx).Infof("订单支付结果通知：%+v", req)

	//查询支付单
	paymentModel := model.OrderPayment{}
	err = l.svcCtx.DbEngine.Model(model.OrderPayment{}).Where("payment_serial_number = ?", req.Data.OrderSerialNumber).First(&paymentModel).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerr.NewBusinessError(xerr.SetCode(xerr.ErrorBusiness), xerr.SetMsg("支付订单不存在"))
	}

	orderRpc := orderclient.NewOrder(zrpc.MustNewClient(l.svcCtx.Config.OrderRpc))
	//查询订单
	orderReply, err := orderRpc.GetOrder(l.ctx, &orderclient.GetOrderReq{OrderSerialNumber: paymentModel.OrderSerialNumber})
	if err != nil {
		return nil, err
	}
	if orderReply.OrderStatus != business.WAIT_PAYMENT_STATE {
		return nil, xerr.NewBusinessError(xerr.SetCode(xerr.ErrorBusiness), xerr.SetMsg("订单不是待支付订单"))
	}

	if paymentModel.ThirdPaymentSerialNumber == "" { //之前没有支付回调过
		err = SuccessPayment(l.ctx, l.svcCtx, orderReply, paymentModel, req)
		if err != nil { //处理失败，可以做其他操作，比如发通知给管理员什么的进入人工处理
			logx.WithContext(l.ctx).Errorf("订单[%s]支付操作失败，通知数据：%+v; 错误：%+v", orderReply.OrderSerialNumber, req, err)
		}
	} else { //之前已有记录支付回调
		if paymentModel.ThirdPaymentSerialNumber == req.Data.PaymentSerialNumber { //重复通知
			//不做任何处理
		} else { //另一个支付通知，属于用户重复支付
			err = RepeatedPayment(l.ctx, l.svcCtx, orderReply, req)
			if err != nil { //处理失败，可以做其他操作，比如发通知给管理员什么的进入人工处理
				logx.WithContext(l.ctx).Errorf("订单[%s]重复支付，做退款处理失败，需要人工介入，通知数据：%+v; 错误：%+v", orderReply.OrderSerialNumber, req, err)
			}
		}
	}
	return &types.PaymentNoticeResp{Code: "success"}, nil
}

// SuccessPayment 支付成功
func SuccessPayment(ctx context.Context, svcCtx *svc.ServiceContext, order *order.GetOrderReply, paymentModel model.OrderPayment, req *types.PaymentNoticeReq) (err error) {
	paymentRpc, err := svcCtx.Config.PaymentRpc.BuildTarget()
	if err != nil {
		return xerr.NewBusinessError(xerr.SetCode(xerr.ErrorBusiness), xerr.SetMsg("支付异常超时"))
	}
	logx.Infof("paymentRpc链接Url：%+v", paymentRpc)
	orderRpc, err := svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return xerr.NewBusinessError(xerr.SetCode(xerr.ErrorBusiness), xerr.SetMsg("支付异常超时"))
	}
	err = dtmdriver.Use("dtm-driver-gozero")
	if err != nil {
		return err
	}
	gid := dtmgrpc.MustGenGid(svcCtx.Config.MicroServiceTarget)
	err = dtmgrpc.TccGlobalTransaction(svcCtx.Config.MicroServiceTarget, gid, func(tcc *dtmgrpc.TccGrpc) error {
		paymentRpcReq := payment.PaymentNoticeReq{
			OrderSerialNumber:   req.Data.OrderSerialNumber,
			PaymentSerialNumber: req.Data.PaymentSerialNumber,
			PaymentAmount:       tool.Float64ToString(req.Data.PaymentAmount, 2),
			PaymentWay:          req.Data.PaymentWay,
			PaymentTime:         req.Data.PaymentTime,
		}
		r := &emptypb.Empty{}
		err = tcc.CallBranch(&paymentRpcReq, paymentRpc+"/payment.payment/orderPaymentNoticeTccTry", paymentRpc+"/payment.payment/orderPaymentNoticeTccConfirm", paymentRpc+"/payment.payment/orderPaymentNoticeTccCancel", r)
		if err != nil {
			return err
		}
		err = tcc.CallBranch(&paymentRpcReq, orderRpc+"/order.order/paymentSuccessOrderTccTry", orderRpc+"/order.order/paymentSuccessOrderTccConfirm", orderRpc+"/order.order/paymentSuccessOrderTccCancel", r)
		return err
	})
	return err

}

// RepeatedPayment 重复支付
func RepeatedPayment(ctx context.Context, svcCtx *svc.ServiceContext, order *order.GetOrderReply, req *types.PaymentNoticeReq) (err error) {
	logx.WithContext(ctx).Slowf("订单[%s]重复支付，做退款处理，通知数据：%+v", order.OrderSerialNumber, req)
	return nil
}
