package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
	"server/game"
)

var skeleton = base.NewSkeleton(game.ChanRPC)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}
