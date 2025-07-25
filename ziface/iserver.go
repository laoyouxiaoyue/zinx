package ziface

type IServer interface {
	Start() // 启动服务器

	Stop() // 停止服务器

	Server() // 运行服务器
	// AddRouter 添加路由
	AddRouter(msgId uint32, router IRouter)
}
