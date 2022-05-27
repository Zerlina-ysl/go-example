package main

import (
	"log"
	"net"
	"net/rpc"
)

//-------------------------------------服务端代码

type HelloService struct{}

func (p *HelloService) Hello(request string,reply *string)error{
	*reply = "hello:"+request
	return nil
}

func main(){
	//RegisterHelloService(svc HelloServiceInterface)
	RegisterHelloService(new(HelloService))
	listener,err:=net.Listen("tcp","localhost:1234")
	if err!=nil{
		log.Fatal("listenTcp error:",err)
	}
	for {
		//为多个tcp连接提供服务
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}


