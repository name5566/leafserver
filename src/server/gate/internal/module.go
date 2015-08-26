package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"server/conf"
	"server/game"
	"server/msg"
)

type Module struct {
	*gate.TCPGate
}

func (m *Module) OnInit() {
	m.TCPGate = &gate.TCPGate{
		Addr:            conf.Server.Addr,
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		LenMsgLen:       conf.LenMsgLen,
		MinMsgLen:       conf.MinMsgLen,
		MaxMsgLen:       conf.MaxMsgLen,
		LittleEndian:    conf.LittleEndian,
		AgentChanRPC:    game.ChanRPC,
	}

	switch conf.Encoding {
	case "json":
		m.TCPGate.JSONProcessor = msg.JSONProcessor
	case "protobuf":
		m.TCPGate.ProtobufProcessor = msg.ProtobufProcessor
	default:
		log.Fatal("unknown encoding: %v", conf.Encoding)
	}
}
