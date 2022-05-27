package main

import (
	"awesomeProject/RPC/HttpRpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main(){
	//基于JSON编码实现RPC
	rpc.RegisterName("HelloService",new(HttpRpc.HelloService))
	listener,err:=net.Listen("tcp","localhost:1234")
	if err!=nil{
		log.Fatal("Listen error:",err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal("accept error:", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}


