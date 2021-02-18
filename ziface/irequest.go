package ziface

type IRequest interface {

	//当前连接
	GetConn() IConnection

	//请求的消息的数据
	GetData() []byte
}


