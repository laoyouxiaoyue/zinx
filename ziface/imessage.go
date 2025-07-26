package ziface

type IMessage interface {
	GetMsgId() uint32 // 获取消息ID

	GetDataLen() uint32 // 获取消息长度

	GetData() []byte // 获取消息内容

	SetMsgId(uint32) // 设置ID

	SetData([]byte) // 设置消息内容

	SetDataLen(uint32) // 设置消息长度
}
