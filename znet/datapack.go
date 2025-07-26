package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/utils"
	"zinx/ziface"
)

// 封包 拆包 处理TCP粘包

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}
func (d *DataPack) GetHeadLen() uint32 {
	// 长度（4）+ID（4）
	return 8
}

func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})

	err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen())
	if err != nil {
		return nil, err
	}
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId())
	if err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

func (d *DataPack) Unpack(data []byte) (ziface.IMessage, error) {
	dataBuff := bytes.NewBuffer(data)

	msg := &Message{}
	err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen)
	if err != nil {
		return nil, err
	}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	//超过最大
	if utils.GlobalObject.MaxPacketSize > 0 && msg.GetDataLen() > uint32(utils.GlobalObject.MaxPacketSize) {
		return nil, errors.New("MessageTooLargeErr")
	}
	return msg, nil
}
