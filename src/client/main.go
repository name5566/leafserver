package main

import (
	"github.com/name5566/leaf/cluster"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/conf"
	"sync"
	"time"
)

type S2S_Test struct {
	ServerName	string
}

func handleTest(args []interface{}) {
	msg := args[0].(*S2S_Test)
	agent := args[1].(*cluster.Agent)
	log.Debug("msgServerName:%v agentServerName:%v", msg.ServerName, agent.ServerName)

	time.Sleep(time.Second)
	sendMsg := &S2S_Test{ServerName:conf.ServerName}
	agent.WriteMsg(sendMsg)
}

func main() {
	conf.ServerName = "client"
	conf.ConnAddrs = []string {"localhost:32017"}
	cluster.Processor.SetHandler(&S2S_Test{}, handleTest)
	cluster.Init()

	time.Sleep(time.Second * 2)
	msg := &S2S_Test{ServerName:conf.ServerName}
	cluster.GetAgent("game").WriteMsg(msg)

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
