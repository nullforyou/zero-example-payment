package logic

import (
	"context"
	"payment/cmd/dao/model"

	"greet-pb/payment/types/payment"
	"payment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderPaymentNoticeTccTryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderPaymentNoticeTccTryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPaymentNoticeTccTryLogic {
	return &OrderPaymentNoticeTccTryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderPaymentNoticeTccTry 支付结果通知
func (l *OrderPaymentNoticeTccTryLogic) OrderPaymentNoticeTccTry(in *payment.PaymentNoticeReq) (*payment.PaymentNoticePayReply, error) {
	logx.WithContext(l.ctx).Info("进入 OrderPaymentNoticeTccTry")
	paymentNotice := model.OrderPaymentNotice{
		PaymentSerialNumber: in.PaymentSerialNumber,
		ThirdNotice:         "OrderPaymentNoticeTccTry",
	}

	l.svcCtx.DbEngine.Create(&paymentNotice)
	return &payment.PaymentNoticePayReply{}, nil
}
