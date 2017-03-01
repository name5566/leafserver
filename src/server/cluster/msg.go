package cluster

import (
	"github.com/name5566/leaf/cluster"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/conf"
)

type S2S_Test struct {
	ServerName	string
}

func handleTest(args []interface{}) {
	msg := args[0].(*S2S_Test)
	agent := args[1].(*cluster.Agent)
	log.Debug("msgServerName:%v agentServerName:%v", msg.ServerName, agent.ServerName)

	sendMsg := &S2S_Test{ServerName:conf.ServerName}
	agent.WriteMsg(sendMsg)
}

func Init() {
	cluster.Processor.SetHandler(&S2S_Test{}, handleTest)
}