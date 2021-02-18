package ziface

import "net"

type IConnection interface {
	//启动链接
	Start()

	// 关闭链接
	Stop()

	//获取当前绑定的 socket conn
	GetTCPConnection() *net.TCPConn

	//获取当前链接的ID
	GetConnID() uint32

	//客户端的Addr
	RemoteAddr() net.Addr

	//发送数据
	Send(data []byte) error
}

//处理链接的方法
type HandleFunc func(*net.TCPConn,[]byte,int) error
