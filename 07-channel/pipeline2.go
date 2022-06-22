package main

import "fmt"
//<-chan int是接收int
//chan int<-发送
func counter(out chan<-int){
	for x:=0;x<100;x++{
		//发送
		out<-x
	}
	close(out)
}
func squarer(out chan <-int,in <-chan int){
	for x:=range in{
	//发送
		out<-x*x
	}
	close(out)
}
func Printer(in <-chan int){
	for x:=range in{
		 fmt.Println(x)
	}
}
func main(){

	naturals := make(chan int)
	squares:= make(chan int)
	//导致naturals隐式转换为发送型
	go counter(naturals)
	go squarer(squares,naturals)
	//任何双向channel向单向channel变量的赋值操作会导致隐式转换
	Printer(squares)
}