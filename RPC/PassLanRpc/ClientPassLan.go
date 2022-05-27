package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main(){
	conn,err:=net.Dial("tcp","localhost:1234")
	if err!=nil{
		log.Fatal("Dial error:",err)
	}
	//基于该链接建立一个针对客户端的JSON编解码器
	//rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	client:=rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply string
	err=client.Call("HelloService.Hello","河大最帅的t",&reply)
	if err!=nil{
		log.Fatal("Call err:",err)
	}
	fmt.Println(reply)
}
