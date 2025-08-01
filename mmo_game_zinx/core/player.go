package core

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log/slog"
	"math/rand"
	"sync"
	"zinx/mmo_game_zinx/pb"
	"zinx/ziface"
)

type Player struct {
	Pid  int32
	Conn ziface.IConnection
	X    float32
	Y    float32
	Z    float32
	V    float32
}

var PidGen int32 = 1
var IdLock sync.Mutex

func NewPlayer(conn ziface.IConnection) *Player {
	IdLock.Lock()
	id := PidGen
	PidGen++
	IdLock.Unlock()

	return &Player{
		Pid:  id,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)),
		Y:    0,
		Z:    float32(134 + rand.Intn(17)),
		V:    0,
	}
}
func (p *Player) SendMsg(msgId uint32, data proto.Message) {
	msg, err := proto.Marshal(data)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	if p.Conn == nil {
		slog.Error("ConnNilErr")
		return
	}
	if err := p.Conn.SendMsg(msgId, msg); err != nil {
		slog.Error(err.Error())
		return
	}
}

func (p *Player) SyncPid() {
	data := &pb.SyncPid{
		Pid: int32(p.Pid),
	}
	p.SendMsg(1, data)
}
func (p *Player) BroadCastStartPosition() {

	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2, //TP2 代表广播坐标
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	p.SendMsg(200, msg)
}

func (p *Player) Talk(content string) {
	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  1,
		Data: &pb.BroadCast_Content{
			Content: content,
		},
	}
	player := WorldManagerObj.GetAllPlayers()
	for _, player := range player {
		player.SendMsg(200, msg)
	}
}

// 给当前玩家周边的(九宫格内)玩家广播自己的位置，让他们显示自己
func (p *Player) SyncSurrounding() {
	//1 根据自己的位置，获取周围九宫格内的玩家pid
	pids := WorldManagerObj.AoiMgr.GetPidsByPos(p.X, p.Z)
	//2 根据pid得到所有玩家对象
	players := make([]*Player, 0, len(pids))
	//3 给这些玩家发送MsgID:200消息，让自己出现在对方视野中
	for _, pid := range pids {
		players = append(players, WorldManagerObj.GetPlayer(int32(pid)))
	}
	//3.1 组建MsgId200 proto数据
	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2, //TP2 代表广播坐标
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	//3.2 每个玩家分别给对应的客户端发送200消息，显示人物
	for _, player := range players {
		player.SendMsg(200, msg)
	}
	//4 让周围九宫格内的玩家出现在自己的视野中
	//4.1 制作Message SyncPlayers 数据
	playersData := make([]*pb.Player, 0, len(players))
	for _, player := range players {
		p := &pb.Player{
			Pid: player.Pid,
			P: &pb.Position{
				X: player.X,
				Y: player.Y,
				Z: player.Z,
				V: player.V,
			},
		}
		playersData = append(playersData, p)
	}

	//4.2 封装SyncPlayer protobuf数据
	SyncPlayersMsg := &pb.SyncPlayers{
		Ps: playersData[:],
	}

	//4.3 给当前玩家发送需要显示周围的全部玩家数据
	p.SendMsg(202, SyncPlayersMsg)
}

func (p *Player) UpdataPos(x float32, y float32, z float32, v float32) {
	p.X = x
	p.Y = y
	p.Z = z
	p.V = v
	slog.Info(fmt.Sprintf("%d :%v,%v,%v,%v", p.Pid, x, y, z, v))
	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  4,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	players := p.GetSurroudingPlayers()
	for _, player := range players {
		slog.Info(fmt.Sprintf("send to player %d", player.Pid))
		player.SendMsg(200, msg)
	}
}

func (p *Player) GetSurroudingPlayers() []*Player {
	pids := WorldManagerObj.AoiMgr.GetPidsByPos(p.X, p.Y)
	player := make([]*Player, 0)
	for _, pid := range pids {
		player = append(player, WorldManagerObj.GetPlayer(int32(pid)))
	}
	return player
}
