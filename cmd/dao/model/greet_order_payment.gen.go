// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameOrderPayment = "greet_order_payment"

// OrderPayment mapped from table <greet_order_payment>
type OrderPayment struct {
	ID                       int64          `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	OrderID                  int64          `gorm:"column:order_id;type:bigint;not null;comment:订单id" json:"order_id"`
	MemberID                 int64          `gorm:"column:member_id;type:bigint;not null;comment:用户id" json:"member_id"`
	OrderSerialNumber        string         `gorm:"column:order_serial_number;type:varchar(30);not null;comment:订单编号" json:"order_serial_number"`
	PaymentStatus            int32          `gorm:"column:payment_status;type:smallint;not null;comment:支付单状态1：已创建待支付；2：已支付；3：支付失败；4:支付异常,需要人工介入;" json:"payment_status"`
	PaymentSerialNumber      string         `gorm:"column:payment_serial_number;type:varchar(32);not null;comment:支付订单号" json:"payment_serial_number"`
	PaymentAmount            float64       `gorm:"column:payment_amount;type:decimal(10,2);comment:支付单支付金额;default:null;" json:"payment_amount"`
	PaymentType              int32          `gorm:"column:payment_type;type:tinyint;not null;comment:支付方式： 1：银联支付 2：支付宝支付 3：微信支付" json:"payment_type"`
	ThirdPaymentSerialNumber string         `gorm:"column:third_payment_serial_number;type:varchar(32);not null;comment:支付订单号" json:"third_payment_serial_number"`
	ThirdPaymentTime         time.Time     `gorm:"column:third_payment_time;type:datetime;comment:第三方通知支付时间;default:null;" json:"third_payment_time"`
	DeletedAt                gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;default:null;" json:"deleted_at"`
	CreatedAt                *time.Time     `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt                *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

// TableName OrderPayment's table name
func (*OrderPayment) TableName() string {
	return TableNameOrderPayment
}