package svc

import (
	"greet/abc/abc"
	"greet/api/internal/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	ABCRpc abc.AbcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	a := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: "dns:///127.0.0.1:8080",
	})
	return &ServiceContext{
		Config: c,
		ABCRpc: abc.NewAbcClient(a.Conn()),
	}
}
