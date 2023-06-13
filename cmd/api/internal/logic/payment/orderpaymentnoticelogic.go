package payment

import (
	"context"

	"payment/cmd/api/internal/svc"
	"payment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return
}
