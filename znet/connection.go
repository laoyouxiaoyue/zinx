package znet

import (
	"errors"
	"fmt"
	"io"
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

	// 告知当前链接已经退出/停止 channel
	ExitChan chan bool

	//消息通道
	MsgChan chan []byte
	// 当前链接处理方法
	MsgHandler ziface.IMsgHandle
}

// NewConnection 初始化链接模块
func NewConnection(conn *net.TCPConn, connID uint32, MsgHandler ziface.IMsgHandle) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: MsgHandler,
		MsgChan:    make(chan []byte, 1024),
		isClosed:   false,
		ExitChan:   make(chan bool),
	}
	return c
}

// StartWriter 写业务
func (c *Connection) StartWriter() {
	slog.Info(fmt.Sprintf("Writer Goroutine is starting [%d]", c.ConnID))
	defer slog.Info(fmt.Sprintf("Writer Goroutine is stopping [%d]", c.ConnID))

	for {
		select {
		case <-c.ExitChan:
			return
		case data := <-c.MsgChan:
			_, err := c.Conn.Write(data)
			if err != nil {
				slog.Error(fmt.Sprintf("Writer Goroutine Write err [%s]", err.Error()))
			}

		}
	}
}

// StartReader 读业务
func (c *Connection) StartReader() {
	slog.Info(fmt.Sprintf("Reader Goroutine is starting [%d]", c.ConnID))
	defer c.Stop()
	defer slog.Info(fmt.Sprintf("Reader Goroutine is stopping [%d]", c.ConnID))
	for {
		//buf := make([]byte, 512)
		//_, err := c.Conn.Read(buf)
		//if err != nil {
		//	slog.Error(fmt.Sprintf("Reader Goroutine [%d]  recv error [%s]", c.ConnID, err))
		//	break
		//}

		dp := NewDataPack()

		headdata := make([]byte, dp.GetHeadLen())

		_, err := io.ReadFull(c.Conn, headdata)
		if err != nil {
			return
		}
		msg, err := dp.Unpack(headdata)
		// 当前request
		if msg.GetDataLen() > 0 {
			data := make([]byte, msg.GetDataLen())
			_, err := io.ReadFull(c.GetTCPConnection(), data)
			if err != nil {
				slog.Error("read data err:", err)
				break
			}
			msg.SetData(data)
		}
		req := &Request{
			conn: c,
			msg:  msg,
		}
		go func(req ziface.IRequest) {
			c.MsgHandler.DoMsgHandler(req)
		}(req)

	}
}
func (c *Connection) Start() {
	slog.Info(fmt.Sprintf("start connection %d", c.ConnID))
	//启动读

	go c.StartReader()
	go c.StartWriter()
}

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("ConnectionClosedErr")
	}
	dp := NewDataPack()

	pack, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		slog.Error("pack err msg id =:", msgId)
		return errors.New("PackMsgErr")
	}
	c.MsgChan <- pack
	return nil
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
	c.ExitChan <- true
	// 回收资源
	close(c.ExitChan)
	close(c.MsgChan)
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
