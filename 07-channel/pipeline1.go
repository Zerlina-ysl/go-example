package main

import "fmt"

func main(){
	naturals := make(chan int)
	squares := make(chan int)

	//counter
	go func() {
		for x:=0;x<100;x++{
			naturals<-x
		}
		//当发送者没有更多的值需要发送，接收者也知道需要停止不必要的接收等待，可以通过close函数
		//当一个被关闭的channels中所有数据接收成功，后续的接收操作不在阻塞而是立即返回零值
		close(naturals)
//向已关闭的channel发送数据将导致panic异常
	}()
	//square
	go func() {
		for{
			for x:=range naturals{
				squares<-x*x
			}
			close(squares)
		}
	}()
	for x:=range squares{
		fmt.Println(x)

	}
	//printer
	//go func() {
	//	for x:=range squares{
	//		fmt.Println(<-squares)
	//
	//	}
	//	}()
}
