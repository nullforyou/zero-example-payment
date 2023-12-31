package payment

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"go-common/tool"
	"greet-pb/payment/paymentclient"
	"greet-pb/payment/types/payment"

	"payment/cmd/api/internal/svc"
	"payment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePaymentLogic {
	return &CreatePaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePaymentLogic) CreatePayment(req *types.PaymentReq) (resp *types.PaymentResp, err error) {
	rpcCreatePaymentReq := payment.CreatePaymentReq{
		OrderSerialNumber:   req.OrderSerialNumber,
		PaymentWay:          req.PaymentWay,
		PaymentAmount:       tool.Float64ToString(req.PaymentAmount, 2),
		TransactionPassword: req.TransactionPassword,
	}

	paymentRpc := paymentclient.NewPayment(zrpc.MustNewClient(l.svcCtx.Config.PaymentRpc))
	rpcRelay, err := paymentRpc.CreatePayment(l.ctx, &rpcCreatePaymentReq)
	if err != nil {
		return nil, err
	}
	return &types.PaymentResp{OrderSerialNumber: rpcRelay.OrderSerialNumber}, err
}
