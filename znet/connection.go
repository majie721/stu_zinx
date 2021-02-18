package znet

import (
	"fmt"
	"mzinx/utils"
	"mzinx/ziface"
	"net"
)

type Connection struct {

	//socket套接字
	Conn *net.TCPConn

	//链接ID
	ConnID uint32

	//连接状态
	IsClose bool


	//告知当前链接已经退出
	ExitChannel chan bool

	//该链接处理的方法 Router
	Router ziface.IRouter
}

func (c *Connection) startReader()  {
	fmt.Printf("conn startReader ... \n")
	defer fmt.Printf("conn startReader  exit..., connid:%d \n",c.ConnID)
	defer c.Stop()


	for {
		buf := make([]byte,utils.GlobalObject.MaxPackageSize)

		_,err :=c.Conn.Read(buf)

		if err != nil{
			fmt.Printf("recv buf read error %v\n",err)
			continue
		}

		req := Request{
			conn: c,
			data: buf,
		}

		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)



	}
}


//启动链接
func (c *Connection) Start()  {
	fmt.Printf("conn start..., connID %d \n",c.ConnID)

	//从当前链接读业务
	go c.startReader()

	//从当前链接写业务

}

// 关闭链接
func (c *Connection) Stop(){
	fmt.Printf("conn stop..., connID %d \n",c.ConnID)

	//判断是否关闭
	if c.IsClose ==true{
		return
	}

	c.IsClose = true

	c.Conn.Close()

	close(c.ExitChannel)
}

//获取当前绑定的 socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn{
	return  c.Conn
}

//获取当前链接的ID
func (c *Connection) GetConnID() uint32{
	return  c.ConnID
}

//客户端的Addr
func (c *Connection) RemoteAddr() net.Addr{
	return c.Conn.RemoteAddr()
}

//发送数据
func (c *Connection) Send(data []byte) error{
	return  nil
}

func NewConnection(conn *net.TCPConn,connID uint32,router ziface.IRouter) *Connection{
	c := &Connection{
		Conn: conn,
		ConnID: connID,
		IsClose: false,
		Router: router,
		ExitChannel: make(chan bool,1),
	}

	return c
}