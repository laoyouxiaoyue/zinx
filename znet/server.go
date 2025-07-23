package znet

import (
	"fmt"
	"log/slog"
	"net"
	"zinx/ziface"
)

type Server struct { // IServer 接口实现 服务器模块
	Name      string // 服务器名称
	IPVersion string // 服务器绑定IP版本
	IP        string // 服务器监听的IP
	Port      int    // 服务器监听的端口
}

func (s *Server) Start() {
	slog.Info(fmt.Sprintf("[Start] Server Listenner at IP :%s,Port %d", s.IP, s.Port))
	// 1 获取TCP Addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		slog.Error("resolve tcp addr err:", err)
		return
	}
	// 2 监听服务器地址
	listenner, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		slog.Error("listen tcp err:", err)
		return
	}
	slog.Info("start server success")
	// 3 阻塞等待客户端连接，处理客户端链接业务

	for {
		conn, err := listenner.AcceptTCP()
		if err != nil {
			slog.Error("accept err:", err)
			continue
		}

		//已经与客户端建立间接
		go func() {
			for {
				buf := make([]byte, 512)
				bufLen, err := conn.Read(buf)
				if err != nil {
					slog.Error("recv buf err:", err)
					continue
				}

				if _, err := conn.Write(buf[:bufLen]); err != nil {
					slog.Error("wirte back buf err:", err)
					continue
				}
			}
		}()
	}
}

func (s *Server) Stop() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Server() {
	s.Start()

	//阻塞状态
	select {}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      80,
	}
	return s
}
