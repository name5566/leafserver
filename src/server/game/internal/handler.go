package internal

import (
	"server/msg"
	"github.com/name5566/leaf/cluster"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/conf"
	"reflect"
)

func handleTest(args []interface{}) {
	recvMsg := args[0].(*msg.S2S_Test)
	agent := args[1].(*cluster.Agent)
	log.Debug("msgServerName:%v agentServerName:%v", recvMsg.ServerName, agent.ServerName)

	sendMsg := &msg.S2S_Test{ServerName:conf.ServerName}
	agent.WriteMsg(sendMsg)
}

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.S2S_Test{}, handleTest)
}