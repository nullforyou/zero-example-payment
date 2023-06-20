package logic

import (
	"context"
	"payment/cmd/dao/model"

	"greet-pb/payment/types/payment"
	"payment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderPaymentNoticeTccConfirmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderPaymentNoticeTccConfirmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPaymentNoticeTccConfirmLogic {
	return &OrderPaymentNoticeTccConfirmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderPaymentNoticeTccConfirmLogic) OrderPaymentNoticeTccConfirm(in *payment.PaymentNoticeReq) (*payment.PaymentNoticePayReply, error) {
	logx.WithContext(l.ctx).Info("进入 OrderPaymentNoticeTccConfirm")

	paymentNotice := model.OrderPaymentNotice{
		PaymentSerialNumber: in.PaymentSerialNumber,
		ThirdNotice:         "OrderPaymentNoticeTccConfirm",
	}

	l.svcCtx.DbEngine.Create(&paymentNotice)

	return &payment.PaymentNoticePayReply{}, nil
}
