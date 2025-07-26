package ziface

// IDataPack 封包 拆包 TCP粘包
type IDataPack interface {
	// GetHeadLen 获取包头长度
	GetHeadLen() uint32
	// Pack 封包
	Pack(msg IMessage) ([]byte, error)
	// UnPack 拆包
	UnPack([]byte) (IMessage, error)
}
