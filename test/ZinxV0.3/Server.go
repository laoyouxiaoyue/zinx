package main

import (
	"zinx/ziface"
	"zinx/znet"
)

// Zinx框架实现服务端测试

type MyRouter struct {
	znet.BaseRouter
}

func (m *MyRouter) Handle(request ziface.IRequest) {
	_, err := request.GetConnection().GetTCPConnection().Write(request.GetData())
	if err != nil {
		return
	}
}
func main() {
	s := znet.NewServer("[zinx V0.1]")
	s.AddRouter(&MyRouter{})
	s.Start()
}
