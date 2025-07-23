package ziface

import "net"

// IConnection 链接模块 抽象接口
type IConnection interface {
	// Start 启动链接
	Start()

	// Stop 停止链接
	Stop()

	// GetTCPConnection 获取链接绑定的socket conn
	GetTCPConnection() *net.TCPConn

	// GetConnID 获取当前链接ID
	GetConnID() uint32

	// RemoteAddr 获取远程客户端TCP状态 IP 端口
	RemoteAddr() net.Addr

	// Send 发送窗口
	Send(date []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
