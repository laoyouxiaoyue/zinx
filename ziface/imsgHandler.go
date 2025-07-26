package ziface

// 消息管理 多路由

type IMsgHandle interface {
	DoMsgHandler(request IRequest)

	AddRouter(msgId uint32, router IRouter)
}
