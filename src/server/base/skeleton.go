package base

import (
	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/module"
	"server/conf"
)

func NewSkeleton(chanRPC *chanrpc.Server) *module.Skeleton {
	skeleton := &module.Skeleton{
		GoLen:              conf.GoLen,
		TimerDispatcherLen: conf.TimerDispatcherLen,
		ChanRPCServer:      chanRPC,
	}
	skeleton.Init()
	return skeleton
}
