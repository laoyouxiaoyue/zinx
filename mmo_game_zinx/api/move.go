package api

import (
	"google.golang.org/protobuf/proto"
	"log/slog"
	"zinx/mmo_game_zinx/core"
	"zinx/mmo_game_zinx/pb"
	"zinx/ziface"
	"zinx/znet"
)

type MoveApi struct {
	znet.BaseRouter
}

func (m *MoveApi) Handle(request ziface.IRequest) {
	msg := &pb.Position{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		slog.Error("PositionUmarshalErr")
		return
	}
	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		slog.Error("GetPropertyErr")
		request.GetConnection().Stop()
		return
	}
	player := core.WorldManagerObj.GetPlayer(pid.(int32))
	player.UpdataPos(msg.X, msg.Y, msg.Z, msg.V)
}
