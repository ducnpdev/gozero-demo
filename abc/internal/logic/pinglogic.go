package logic

import (
	"context"

	"greet/abc/abc"
	"greet/abc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *abc.Request) (*abc.Response, error) {
	// todo: add your logic here and delete this line

	l.Logger.Debugf("rpc ping %s", in.Ping)

	return &abc.Response{
		Pong: string("rpc pong"),
	}, nil
}
