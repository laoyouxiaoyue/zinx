package ziface

// IRequest 把链接信息 和请求数据包装
type IRequest interface {
	// GetConnection 得到当前连接
	GetConnection() IConnection

	// GetData 得到当前数据
	GetData() []byte

	// GetMsgID
	GetMsgID() uint32
}
