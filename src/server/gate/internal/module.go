package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"server/conf"
	"server/game"
	"server/msg"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		TCPAddr:         conf.Server.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		AgentChanRPC:    game.ChanRPC,
	}

	switch conf.Encoding {
	case "json":
		m.Gate.JSONProcessor = msg.JSONProcessor
	case "protobuf":
		m.Gate.ProtobufProcessor = msg.ProtobufProcessor
	default:
		log.Fatal("unknown encoding: %v", conf.Encoding)
	}
}
