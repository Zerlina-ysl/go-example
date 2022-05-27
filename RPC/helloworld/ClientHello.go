package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main(){
	//拨号RPC服务
	client,err:=rpc.Dial("tcp",":1234")
	if err!=nil{
		log.Fatal("dailing error:",err)
	}
	var reply string
	//调用具体的rpc方法
	//第一个参数： 点号链接的rpc服务名字和方法名字
	//第二个参数和第三个参数是定义RPC方法的两个参数 request string,reply *string

	err = client.Call("HelloService.Hello","hello",&reply)
	if err!=nil{
		log.Fatal("call error:",err)
	}
	fmt.Println(reply)
}
