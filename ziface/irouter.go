package ziface

// IRouter 路由抽象接口
type IRouter interface {
	// PreHandle 处理链接业务之前的钩子方法
	PreHandle(request IRequest)
	// Handle 处理业务的主方法
	Handle(request IRequest)
	// PostHandle 处理业务之后的方法
	PostHandle(request IRequest)
}
