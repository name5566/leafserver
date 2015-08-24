package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
	"server/login"
)

func init() {
	login.Module = new(Module)
}

var skeleton = base.NewSkeleton(login.ChanRPC)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}
