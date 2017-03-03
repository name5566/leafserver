package cluster

import (
	"server/game"
	"server/msg"
	"github.com/name5566/leaf/cluster"
)

func Init()  {
	cluster.Processor.SetRouter(&msg.S2S_Test{}, game.ChanRPC)

	cluster.SetRoute("f0", game.ChanRPC)
	cluster.SetRoute("f1", game.ChanRPC)
	cluster.SetRoute("fn", game.ChanRPC)
	cluster.SetRoute("add", game.ChanRPC)
}
