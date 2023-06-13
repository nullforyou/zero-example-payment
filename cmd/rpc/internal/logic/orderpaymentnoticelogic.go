package logic

import (
	"context"
	"github.com/dtm-labs/client/dtmgrpc"
	"go-common/tool"
	"greet-pb/payment/types/payment"
	"payment/cmd/rpc/internal/process/tcc"
	"payment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderPaymentNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderPaymentNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPaymentNoticeLogic {
	return &OrderPaymentNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderPaymentNotice 支付结果通知
func (l *OrderPaymentNoticeLogic) OrderPaymentNotice(in *payment.PaymentNoticeReq) (*payment.PaymentNoticePayReply, error) {

	transReq := &tcc.TransReq{
		OrderSerialNumber:   in.OrderSerialNumber,
		PaymentSerialNumber: in.PaymentSerialNumber,
		PaymentWay:          in.PaymentWay,
		PaymentAmount:       tool.StringToFloat64(in.PaymentAmount),
		PaymentTime:         in.PaymentTime,
	}

	gid := dtmgrpc.MustGenGid(l.svcCtx.Config.MicroService.Target)
	err := dtmgrpc.TccGlobalTransaction(l.svcCtx.Config.MicroService.Target, gid, func(tcc *dtmgrpc.TccGrpc) error {
		return tcc.CallBranch(transReq, "", "")
	})
	return &payment.PaymentNoticePayReply{}, nil
}
