package znet

import (
	"log/slog"
	"sync"
	"zinx/ziface"
)

type ConnManager struct {
	connections map[uint32]ziface.IConnection
	connLock    sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
		connLock:    sync.RWMutex{},
	}
}
func (c *ConnManager) Add(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	c.connections[conn.GetConnID()] = conn
	slog.Info("add connection %d success", conn.GetConnID())
}

func (c *ConnManager) Remove(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	delete(c.connections, conn.GetConnID())
	slog.Info("remove connection %d success", conn.GetConnID())

}

func (c *ConnManager) Get(connID uint32) ziface.IConnection {
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	if conn, ok := c.connections[connID]; ok {
		return conn
	} else {
		slog.Info("connManager Get connID:%d not exist", connID)
		return nil
	}
}

func (c *ConnManager) Len() int {
	return len(c.connections)
}

func (c *ConnManager) ClearConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	for connId, conn := range c.connections {
		conn.Stop()
		delete(c.connections, connId)
		slog.Info("clear connection %d success", connId)
	}
}
