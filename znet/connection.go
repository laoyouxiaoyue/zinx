package znet

import (
	"fmt"
	"log/slog"
	"net"
	"zinx/ziface"
)

type Connection struct {

	// 链接的socket TCP套接字
	Conn *net.TCPConn

	// 链接ID
	ConnID uint32

	// 链接状态
	isClosed bool

	// 当前链接所绑定处理业务方法API
	handleAPI ziface.HandleFunc

	// 告知当前链接已经退出/停止 channel
	ExitChan chan bool
}

// NewConnection 初始化链接模块
func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleAPI: callback_api,
		isClosed:  false,
		ExitChan:  make(chan bool),
	}
	return c
}

// 读业务
func (c *Connection) StartReader() {
	slog.Info(fmt.Sprintf("Reader Goroutine is starting [%d]", c.ConnID))
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			slog.Error(fmt.Sprintf("Reader Goroutine [%d]  recv error [%s]", c.ConnID, err))
			break
		}
		err = c.handleAPI(c.Conn, buf, cnt)
		if err != nil {
			slog.Error(fmt.Sprintf("Conn ID [%d]  Handle error [%s]", c.ConnID, err))
			break
		}
	}
}
func (c *Connection) Start() {
	slog.Info(fmt.Sprintf("start connection %d", c.ConnID))
	//启动读

	go c.StartReader()
}

func (c *Connection) Stop() {
	slog.Info(fmt.Sprint("Connection Stop ", c.ConnID))

	if c.isClosed {
		return
	}
	c.isClosed = true

	// 关闭
	err := c.Conn.Close()
	if err != nil {
		return
	}

	// 回收资源
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(date []byte) error {
	_, err := c.Conn.Write(date)
	if err != nil {
		return err
	}
	return nil
}
