package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-base/custom_validate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"greet-pb/payment/paymentclient"
	"payment/cmd/api/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	DbEngine   *gorm.DB
	PaymentRpc paymentclient.Payment
	Validator  custom_validate.Validator
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "greet_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,     // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:     c,
		DbEngine:   db,
		Validator:  custom_validate.InitValidator(),
		PaymentRpc: paymentclient.NewPayment(zrpc.MustNewClient(c.PaymentRpc)),
	}
}
