package main

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

