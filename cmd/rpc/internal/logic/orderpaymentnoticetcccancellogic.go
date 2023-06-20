package logic

import (
	"context"
	"payment/cmd/dao/model"

	"greet-pb/payment/types/payment"
	"payment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderPaymentNoticeTccCancelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderPaymentNoticeTccCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPaymentNoticeTccCancelLogic {
	return &OrderPaymentNoticeTccCancelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderPaymentNoticeTccCancelLogic) OrderPaymentNoticeTccCancel(in *payment.PaymentNoticeReq) (*payment.PaymentNoticePayReply, error) {
	logx.WithContext(l.ctx).Info("进入 OrderPaymentNoticeTccCancel")
	paymentNotice := model.OrderPaymentNotice{
		PaymentSerialNumber: in.PaymentSerialNumber,
		ThirdNotice:         "OrderPaymentNoticeTccCancel",
	}

	l.svcCtx.DbEngine.Create(&paymentNotice)
	return &payment.PaymentNoticePayReply{}, nil
}
