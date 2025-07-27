package znet

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

type Server struct { // IServer 接口实现 服务器模块
	Name        string                        // 服务器名称
	IPVersion   string                        // 服务器绑定IP版本
	IP          string                        // 服务器监听的IP
	Port        int                           // 服务器监听的端口
	Router      ziface.IRouter                // 添加Router
	MsgHandler  ziface.IMsgHandle             // 多路由管理
	ConnManager ziface.IConnManager           //链接管理
	OnConnStart func(conn ziface.IConnection) // 启动HOOK函数
	OnConnStop  func(conn ziface.IConnection) // 停止HOOK函数
}

func CallBackToClient(conn *net.TCPConn, date []byte, cnt int) error {
	// 回显业务
	_, err := conn.Write(date[:cnt])
	if err != nil {
		slog.Error(fmt.Sprintf("call back to client error: %v", err))
		return errors.New("CallBackToClientError")
	}
	return nil
}
func (s *Server) Start() {
	slog.Info(fmt.Sprintf("[Start] Server Listenner at IP :%s,Port %d", s.IP, s.Port))
	// 1 获取TCP Addr

	//开启消息队列
	s.MsgHandler.StartWorkerPool()

	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		slog.Error("resolve tcp addr err")
		return
	}
	// 2 监听服务器地址
	listenner, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		slog.Error("listen tcp err")
		return
	}
	slog.Info("start server success")
	// 3 阻塞等待客户端连接，处理客户端链接业务
	var cid uint32
	for {
		conn, err := listenner.AcceptTCP()
		if err != nil {
			slog.Error("accept err:")
			continue
		}
		if s.ConnManager.Len() >= utils.GlobalObject.MaxConn {
			conn.Close()
			continue
		}

		//已经与客户端建立间接 处理新连接的业务方法
		dealConn := NewConnection(s, conn, cid, s.MsgHandler)
		cid++

		go dealConn.Start()
	}

}

func (s *Server) GetConnManager() ziface.IConnManager {
	return s.ConnManager

}
func (s *Server) Stop() {
	slog.Info("stop server success")
	s.ConnManager.ClearConn()
}

func (s *Server) Server() {
	s.Start()

	//阻塞状态
	select {}

}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
}
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:        name,
		IPVersion:   "tcp4",
		IP:          "0.0.0.0",
		Port:        7777,
		MsgHandler:  NewMsgHandle(),
		ConnManager: NewConnManager(),
	}
	return s
}

func (s *Server) SetOnConnStart(f func(connection ziface.IConnection)) {
	s.OnConnStart = f
}
func (s *Server) SetOnConnStop(f func(connection ziface.IConnection)) {
	s.OnConnStop = f
}
func (s *Server) CallOnConnStart(connection ziface.IConnection) {
	if s.OnConnStart != nil {
		s.OnConnStart(connection)
	}

}
func (s *Server) CallOnConnStop(connection ziface.IConnection) {
	if s.OnConnStop != nil {
		s.OnConnStop(connection)
	}
}
