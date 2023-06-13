package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource  string
		TablePrefix string
	}
	OrderRpc     zrpc.RpcClientConf
	MicroService struct {
		Driver   string
		Target   string
		EndPoint string
	}
}
