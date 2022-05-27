package main

import (
	"log"
	"net"
	"net/rpc"
)

//定义HelloService类型

type HelloService struct{}

//其中Hello方法必须
//1 只能有两个可序列化的参数，第二个参数是指针类型
//2 返回error类型
//3 方法是公开方法

func (p *HelloService)Hello(request string,reply *string)error{
	*reply = "hello:"+request
	return nil

}
func main(){
	//将所有满足rpc规则的对象类型的对象方法注册为rpc函数
	//所有注册的方法都会放在HelloService服务的空间下
	//内建new函数创建HelloService结构体变量
	rpc.RegisterName("HelloService",new(HelloService))
	listener,err:= net.Listen("tcp",":1234")
	if err!=nil{
		log.Fatal("ListenTcp error:",err)
	}
	conn,err:=listener.Accept()
	if err!=nil{
		log.Fatal("Accept error:",err)
	}
	//在该Tcp链接上为对方提供服务
	rpc.ServeConn(conn)

}