package main

import (
	"fmt"
	"log/slog"
	"zinx/mmo_game_zinx/api"
	"zinx/mmo_game_zinx/core"
	"zinx/ziface"
	"zinx/znet"
)

func OnConnectionAdd(conn ziface.IConnection) {
	player := core.NewPlayer(conn)
	player.SyncPid()
	player.BroadCastStartPosition()
	core.WorldManagerObj.AddPlayer(player)
	conn.SetProperty("pid", player.Pid)
	player.SyncSurrounding()
	slog.Info(fmt.Sprintf("player %d join the game", player.Pid))
}

func main() {
	s := znet.NewServer("MMO GAME")
	s.SetOnConnStart(OnConnectionAdd)
	s.AddRouter(2, &api.WorldChatApi{})
	s.AddRouter(3, &api.MoveApi{})
	s.Server()
}
