package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"server/conf"
	game "server/game/internal"
	gate "server/gate/internal"
	login "server/login/internal"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath

	leaf.Run(
		new(game.Module),
		new(gate.Module),
		new(login.Module),
	)
}
