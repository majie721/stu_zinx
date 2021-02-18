package ziface

type IServer interface {

	//启动
	Start()

	//运行
	Serve()

	//停止
	Stop()

	//注册路由方法
	AddRouter(router IRouter)
}