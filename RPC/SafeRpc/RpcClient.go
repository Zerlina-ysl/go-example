package main

import (
	"fmt"
	"log"
)

//根据规范编写rpc调用代码
func main(){
	var client, err = DialHelloService("tcp", "localhost:1234")
	if err!=nil{
		log.Fatal("dial error:",err)
	}
	var reply string
	//直接通过接口对应的方法调用rpc函数
	err=client.Hello("河大最帅的t",&reply)
	if err!=nil{
		log.Fatal("call error:",err)
	}
	fmt.Println(reply)
}
