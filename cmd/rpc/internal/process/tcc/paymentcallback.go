package tcc

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"payment/cmd/dao/model"
	"payment/cmd/rpc/internal/svc"
)

type TransReq struct {
	OrderSerialNumber   string  //订单编号
	PaymentSerialNumber string  //支付单编号
	PaymentWay          int64   //支付方式 1：银联支付 2：支付宝支付 3：微信支付
	PaymentAmount       float64 //支付金额
	PaymentTime         int64   //支付时间戳
}

func (r TransReq) ProtoReflect() protoreflect.Message {
	//TODO implement me
	panic("implement me")
}

func PaymentSuccessTry(svcCtx *svc.ServiceContext, req TransReq) {
	noticeModel := model.OrderPaymentNotice{PaymentSerialNumber: req.PaymentSerialNumber}
	noticeModel.ThirdNotice = "{'title': '这是 PaymentSuccessTry'}"
	svcCtx.DbEngine.Create(&noticeModel)
}

func PaymentSuccessConfirm(svcCtx *svc.ServiceContext, req TransReq) {
	noticeModel := model.OrderPaymentNotice{PaymentSerialNumber: req.PaymentSerialNumber}
	noticeModel.ThirdNotice = "{'title': '这是 PaymentSuccessConfirm'}"
	svcCtx.DbEngine.Create(&noticeModel)
}
func PaymentSuccessCancel(svcCtx *svc.ServiceContext, req TransReq) {
	noticeModel := model.OrderPaymentNotice{PaymentSerialNumber: req.PaymentSerialNumber}
	noticeModel.ThirdNotice = "{'title': '这是 PaymentSuccessCancel'}"
	svcCtx.DbEngine.Create(&noticeModel)
}
