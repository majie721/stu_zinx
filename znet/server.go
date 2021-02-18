package znet

import (
	"fmt"
	"mzinx/utils"
	"mzinx/ziface"
	"net"
)

type Server struct {
	//服务name
	Name string

	//绑定的ip版本
	IPVersion string

	//监听ip
	IP string

	//监听的端口
	Port int

	//当前的Server 添加router
	Router ziface.IRouter
}


func (s *Server) Start() {
	fmt.Printf("%#v",utils.GlobalObject)
	fmt.Printf("[start] server listenner at  %s:%d \n", s.IP, s.Port)
	go func() {
		//1.创建套接字
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve ip addr error:", err)
			return
		}
		//2.listen 服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen tcp error:", err)
			return
		}

		var cid uint32 =0

		fmt.Printf("[start] server listener at  %s:%d success \n", s.IP, s.Port)
		//3.阻塞可速断链接,处理客户端业务(读写)
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept ERROR ... ")
				continue
			}


			dealConn := NewConnection(conn,cid,s.Router)
			cid++

			//启动当前链接处理
			go dealConn.Start()
		}
	}()
}

func (s *Server) Serve() {

	//启动服务
	s.Start()

	//阻塞状态
	select {
	}
}

func (s *Server) Stop() {
	//todo
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router  = router
	fmt.Println("add router success")
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TCPPort,
		Router: nil,
	}

	return s
}

