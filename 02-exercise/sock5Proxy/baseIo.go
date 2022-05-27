package main

import (
	"bufio"

	"log"
	"net"
)
//tcp echo server
//返回输入信息的tcp server
func main(){
	//Listen监听一个网络端口上的连接
	server,err:=net.Listen("tcp","127.0.0.1:1080")
	if err!=nil{
		panic(err)
	}
	for{
		//Accept()会阻塞 直到新的连接被创建 并返回一个net.Conn来表示该连接
		client,err:=server.Accept()
		if err!=nil{
			log.Printf("Accept Failed %v",err)
			continue
		}
		//加入go关键字 每一次handleConn的调用都进入一个独立的goroutine 使其支持并发
		go process1(client)
	}
}
func process1(conn net.Conn){
	//defer语句用于延迟一个函数或方法的执行，会在外围函数或者方法返回之前但是返回值计算之后执行，可以在一个延迟执行的函数内部修改函数的命名返回值
	//如果一个函数有多个defer语句，以LIFO的顺序执行
	//该模式创建了一个值，在垃圾收集之前延迟执行一些关闭函数来清理该值
	defer conn.Close()
	//创建缓冲流
	reader:=bufio.NewReader(conn)
	for{
		//每次读一个字节
		b,err:=reader.ReadByte()
		if err!=nil{
			break
		}
		//将字节写入
		_,err=conn.Write([]byte{b})
		if err!=nil{
			break
		}
	}


}
