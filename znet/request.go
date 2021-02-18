package znet

import "mzinx/ziface"

type Request struct {
	//与客户端建立好的链接
	conn ziface.IConnection

	//客户端数据
	data []byte
}

//当前连接
func (r *Request)GetConn() ziface.IConnection  {
	return  r.conn
}

//请求的消息的数据
func (r *Request) GetData() []byte{
	return  r.data
}
