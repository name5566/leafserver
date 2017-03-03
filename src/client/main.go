package main

import (
	"github.com/name5566/leaf/cluster"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/conf"
	"time"
	"server/base"
	"fmt"
	"github.com/name5566/leaf/chanrpc"
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

func rpcTest(agent *cluster.Agent, chanAsynRet chan *chanrpc.RetInfo)  {
	// sync
	err := agent.Call0("f0")
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second)

	r1, err := agent.Call1("f1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r1)
	}
	time.Sleep(time.Second)

	rn, err := agent.CallN("fn")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rn[0], rn[1], rn[2])
	}
	time.Sleep(time.Second)

	ra, err := agent.Call1("add", 1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ra)
	}
	time.Sleep(time.Second)

	// asyn
	agent.AsynCall(chanAsynRet, "f0", func(err error) {
		if err != nil {
			fmt.Println(err)
		}
	})
	time.Sleep(time.Second)

	agent.AsynCall(chanAsynRet, "f1", func(ret interface{}, err error) {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ret)
		}
	})
	time.Sleep(time.Second)

	agent.AsynCall(chanAsynRet, "fn", func(ret []interface{}, err error) {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ret[0], ret[1], ret[2])
		}
	})
	time.Sleep(time.Second)

	agent.AsynCall(chanAsynRet, "add", 1, 2, func(ret interface{}, err error) {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ret)
		}
	})
	time.Sleep(time.Second)

	// go
	agent.Go("f0")
}

func main() {
	conf.ServerName = "client"
	conf.ConnAddrs = []string {"localhost:32017"}
	cluster.Processor.SetHandler(&S2S_Test{}, handleTest)
	cluster.Init()

	time.Sleep(time.Second * 2)
	agent := cluster.GetAgent("game")

	closeSig := make(chan bool)
	skeleton := base.NewSkeleton()

	go rpcTest(agent, skeleton.GetChanAsynRet())
	skeleton.Run(closeSig)
}
