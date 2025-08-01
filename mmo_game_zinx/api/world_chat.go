package api

import (
	"google.golang.org/protobuf/proto"
	"log/slog"
	"zinx/mmo_game_zinx/core"
	"zinx/mmo_game_zinx/pb"
	"zinx/ziface"
	"zinx/znet"
)

type WorldChatApi struct {
	znet.BaseRouter
}

func (this *WorldChatApi) Handle(request ziface.IRequest) {
	msg := &pb.Talk{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		slog.Error("TalkUnmarshalErr")
		return
	}

	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		slog.Error("GetPropertyErr", err)
		request.GetConnection().Stop()
		return
	}
	player := core.WorldManagerObj.GetPlayer(pid.(int32))
	player.Talk(msg.Content)
}
