package znet

import (
	"fmt"
	"zinx/utils"
	"zinx/ziface"
)

type MsgHandle struct {
	// id对应Router处理方法
	Apis map[uint32]ziface.IRouter
	// 负责worker取任务的消息队列
	TaskQueue []chan ziface.IRequest
	// 业务工作worker池的worker数量
	WorkerPoolSize uint32
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
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

// StartWorkerPool 启动工作池
func (m *MsgHandle) StartWorkerPool() {
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		m.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		go m.StartOneWorker(i)
	}
}

func (m *MsgHandle) StartOneWorker(num int) {
	for {
		select {
		case request := <-m.TaskQueue[num]:
			m.DoMsgHandler(request)
		}
	}

}

// SendMsgToTaskQueue 将消息交给WorkerPool
func (m *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	workerID := request.GetConnection().GetConnID()
	m.TaskQueue[workerID%m.WorkerPoolSize] <- request
}
