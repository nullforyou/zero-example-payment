// Code generated by goctl. DO NOT EDIT.
package types

type PaymentReq struct {
	OrderSerialNumber   string  `path:"order_serial_number"`
	PaymentWay          int64   `json:"payment_way" validate:"min=1,max=10"`
	PaymentAmount       float64 `json:"payment_amount" validate:"gt=0"`
	TransactionPassword string  `json:"transaction_password" validate:"min=6"`
}

type PaymentResp struct {
	OrderSerialNumber string `path:"order_serial_number"`
}

type PaymentNoticeReq struct {
	Code int64                   `json:"code"`
	Data PaymentNoticePayloadReq `json:"data"`
}

type PaymentNoticePayloadReq struct {
	OrderSerialNumber   string  `json:"order_serial_number"`
	PaymentWay          int64   `json:"payment_way"`
	PaymentAmount       float64 `json:"payment_amount"`
	PaymentSerialNumber string  `json:"payment_serial_number"`
	PaymentTime         int64   `json:"payment_time"`
}

type PaymentNoticeResp struct {
	Code string `json:"code"`
}