package main

import (
	"net/rpc"
)
//1.服务的名字（增加包路径前缀避免名字冲突

const HelloServiceName = "path/to/pkg.HelloService"

//2.服务要实现的详细方法列表

type HelloServiceInterface = interface {
	Hello(request string,reply *string) error
}
//3.注册该类型服务的函数 传入的对象需满足HelloServiceInterface接口

func RegisterHelloService(svc HelloServiceInterface)error{
	return rpc.RegisterName(HelloServiceName,svc)
}


//----------------------------------------------客户端规范
// 在rpc规范中对客户端调用进行封装

type HelloServiceClient struct{
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

//直接拨号HelloService

func DialHelloService(network,address string)(*HelloServiceClient,error){
	c,err:=rpc.Dial(network,address)
	if err!=nil{
		return nil,err
	}
	return &HelloServiceClient{Client:c},nil
}
func (p *HelloServiceClient)Hello(request string,reply *string)error{
	//服务名称+服务方法 服务方法Hello()的两个参数
	return p.Client.Call(HelloServiceName+".Hello",request,reply)
}
