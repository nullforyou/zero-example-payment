package logic

import (
	"context"
	"go-common/tool"
	"go-zero-base/utils/xerr"
	"greet-pb/order/orderclient"
	"payment/cmd/business"
	"payment/cmd/dao/model"
	"time"

	"greet-pb/payment/types/payment"
	"payment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePaymentLogic {
	return &CreatePaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreatePayment 创建支付单
func (l *CreatePaymentLogic) CreatePayment(in *payment.CreatePaymentReq) (*payment.CreatePaymentReply, error) {

	//TODO 验证密码

	//查询订单
	order, err := l.svcCtx.OrderRpc.GetOrder(l.ctx, &orderclient.GetOrderReq{OrderSerialNumber: in.OrderSerialNumber})
	if err != nil {
		return nil, err
	}
	if order.OrderStatus != business.WAIT_PAYMENT_STATE {
		return nil, xerr.NewBusinessError(xerr.SetCode(xerr.ErrorBusiness), xerr.SetMsg("订单不是待支付订单"))
	}

	//创建支付单
	var paymentModel = model.OrderPayment{
		OrderID:             order.ID,
		MemberID:            order.MemberID,
		OrderSerialNumber:   order.OrderSerialNumber,
		PaymentStatus:       1,
		PaymentSerialNumber: business.GenerateOrderNumber(time.Now()),
		PaymentAmount:       tool.StringToFloat64(in.GetPaymentAmount()),
		PaymentType:         int32(in.PaymentWay),
	}

	l.svcCtx.DbEngine.Save(&paymentModel)

	return &payment.CreatePaymentReply{OrderSerialNumber: order.OrderSerialNumber, PaymentSerialNumber: paymentModel.PaymentSerialNumber}, nil
}
