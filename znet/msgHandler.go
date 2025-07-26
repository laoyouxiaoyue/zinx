package znet

import (
	"fmt"
	"zinx/ziface"
)

type MsgHandle struct {
	// id对应Router处理方法
	Apis map[uint32]ziface.IRouter
}  

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}
func (m *MsgHandle) DoMsgHandler(request ziface.IRequest) {

	router, ok := m.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgID =", request.GetMsgID(), "Not Found")
	}
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}

func (m *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := m.Apis[msgId]; ok {
		panic(fmt.Sprintf("[msgId:%d] already exist", msgId))
	}
	m.Apis[msgId] = router
}
