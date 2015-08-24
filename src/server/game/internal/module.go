package internal

import (
	"server/base"
	"server/game"
)

func init() {
	game.Module = new(Module)
}

var skeleton = base.NewSkeleton(game.ChanRPC)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}
