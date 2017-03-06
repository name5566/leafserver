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

func rpcTest(chanAsynRet chan *chanrpc.RetInfo)  {
	// sync
	err := cluster.Call0("game", "f0")
	if err != nil {
		fmt.Println(err)
	}

	r1, err := cluster.Call1("game", "f1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r1)
	}

	rn, err := cluster.CallN("game", "fn")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rn[0], rn[1], rn[2])
	}

	ra, err := cluster.Call1("game", "add", 1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ra)
	}

	// asyn
	cluster.AsynCall("game", chanAsynRet, "f0", func(err error) {
		if err != nil {
			fmt.Println(err)
		}
	})

	cluster.AsynCall("game", chanAsynRet, "f1", func(ret interface{}, err error) {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ret)
		}
	})

	cluster.AsynCall("game", chanAsynRet, "fn", func(ret []interface{}, err error) {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ret[0], ret[1], ret[2])
		}
	})

	cluster.AsynCall("game", chanAsynRet, "add", 1, 2, func(ret interface{}, err error) {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ret)
		}
	})

	// go
	cluster.Go("game", "f0")
}

func main() {
	conf.ServerName = "client"
	conf.ConnAddrs = []string {"localhost:32017"}
	cluster.Processor.SetHandler(&S2S_Test{}, handleTest)
	cluster.Init()

	time.Sleep(time.Second * 2)

	closeSig := make(chan bool)
	skeleton := base.NewSkeleton()

	go rpcTest(skeleton.GetChanAsynRet())
	skeleton.Run(closeSig)
}
