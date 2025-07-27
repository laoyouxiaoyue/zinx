package ziface

type IConnManager interface {
	// Add 添加连接
	Add(conn IConnection)
	// Remove 删除链接
	Remove(conn IConnection)
	// Get 获取链接
	Get(connID uint32) IConnection
	// Len 连接总数
	Len() int
	// ClearConn 清除所有
	ClearConn()
}
