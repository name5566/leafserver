package game

import (
	"github.com/name5566/leaf/chanrpc"
	"server/conf"
)

var ChanRPC = chanrpc.NewServer(conf.ChanRPCLen)
