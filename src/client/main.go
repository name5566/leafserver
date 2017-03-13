package main

import (
	"github.com/name5566/leaf/cluster"
	"github.com/name5566/leaf/log"
	lconf "github.com/name5566/leaf/conf"
	"client/conf"
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
	sendMsg := &S2S_Test{ServerName:lconf.ServerName}
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

	printRequestCount := func() {
		fmt.Printf("request count %v\n", cluster.GetRequestCount())
	}

	// asyn
	cluster.AsynCall("game", chanAsynRet, "f0", func(err error) {
		if err != nil {
			fmt.Println(err)
		}
	})
	printRequestCount()

	cluster.AsynCall("game", chanAsynRet, "f1", func(ret interface{}, err error) {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ret)
		}
	})
	printRequestCount()

	cluster.AsynCall("game", chanAsynRet, "fn", func(ret []interface{}, err error) {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ret[0], ret[1], ret[2])
		}
	})
	printRequestCount()

	cluster.AsynCall("game", chanAsynRet, "add", 1, 2, func(ret interface{}, err error) {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ret)
		}
	})
	printRequestCount()

	// go
	cluster.Go("game", "f0")
	printRequestCount()

	time.Sleep(time.Second)
	printRequestCount()

	qpsTest := func() {
		for {
			cluster.Call1("game", "qpsTest")
		}

	}
	for i := 0; i < 1000; i++ {
		go qpsTest()
	}
}

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath
	lconf.ServerName = conf.Server.ServerName
	lconf.ListenAddr = conf.Server.ListenAddr
	lconf.ConnAddrs = conf.Server.ConnAddrs
	lconf.PendingWriteNum = conf.Server.PendingWriteNum

	//conf.HeartBeatInterval = 11 //故意让心跳太久，使得game server能自动测试并断线重连
	cluster.Processor.Register(&S2S_Test{})
	cluster.Processor.SetHandler(&S2S_Test{}, handleTest)
	cluster.Init()

	time.Sleep(time.Second * 2)

	closeSig := make(chan bool)
	skeleton := base.NewSkeleton()

	go rpcTest(skeleton.GetChanAsynRet())
	skeleton.Run(closeSig)
}
