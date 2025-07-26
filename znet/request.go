package znet

import "zinx/ziface"

type Request struct {
	// 当前链接
	conn ziface.IConnection

	// 当前数据
	msg ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
