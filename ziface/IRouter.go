package ziface

/**
路由的抽象接口
路由的数据都是IRequest
*/

type IRouter interface {
	//处理conn之前的方法hook
	PreHandle(request IRequest)

	//处理conn方法
	Handle(request IRequest)

	//处理conn之后的方法hook
	PostHandle(request IRequest)
}


