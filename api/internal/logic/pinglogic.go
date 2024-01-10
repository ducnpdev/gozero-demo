package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"greet/abc/abc"
	"greet/api/internal/svc"
	"greet/api/internal/types"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (resp *types.Resp, err error) {

	resp = new(types.Resp)
	// l.Infof("a %s", "123")
	resp.Msg = "pong"
	// l.Info("xin chao day dev test")
	// l.Info("xin chao day dev test")
	// l.Info("xin chao day dev test")
	// l.Info("xin chao day dev test")
	// l.Info("xin chao day dev test")

	l.testOther(l.ctx)

	return
}

func (l *PingLogic) testOther(ctx context.Context) error {
	// l.Info("testOther")

	_, b := l.svcCtx.ABCRpc.Ping(l.ctx, &abc.Request{
		Ping: "xin chao, rest api",
	})
	if b != nil {
		l.Error(b)
	}
	return nil
}

func (l *PingLogic) Ping123() (resp *types.Resp, err error) {
	resp = new(types.Resp)
	resp.Msg = "pong123"
	l.svcCtx.ABCRpc.Ping(l.ctx, &abc.Request{
		Ping: "xin chao, rest api",
	})
	return
}
