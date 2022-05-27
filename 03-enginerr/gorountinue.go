package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	fmt.Println("=== Run TestManyGo")
	//最初尝试goroutinue
	//HelloGoRoutine()
	//通过通信来共享进程
	//CalSquare()
	//通过共享进程来通信
	//Add()
	//使用WaitGroup优化HelloGoRoutine
	ManyGoWait()
}

func Hello(i int){
	println("hello goroutine:"+fmt.Sprint(i))
}
func HelloGoRoutine(){
	for i:=0;i<5;i++ {
		go func(j int){
			Hello(j)
		}(i)
	}
	time.Sleep(time.Second)
}
func CalSquare(){
	//无缓存通道
	src:=make(chan int)
	//有缓冲通道
	dest:=make(chan int,3)
	go func(){
		defer close(src)

		for i:=0;i<10;i++{
			src<-i
		}
	}()
	go func() {
		defer close(dest)
		for i:=range src{
			dest<-i*i
		}
	}()
	for i:=range dest{
		println(i)
	}
}
//通过共享内存进行通信
var(
	x int64
	lock sync.Mutex
)
func addWithLock(){
	for i:=0;i<2000;i++{
		lock.Lock()
		x+=1
		lock.Unlock()
	}
}
func addWithoutLock(){
	for i:=0;i<2000;i++{
		x+=1
	}
}
func Add(){
	x=0
	for i:=0;i<5;i++{
		go addWithoutLock()
	}
	time.Sleep(time.Second)
	println("withoutlock",x)
	x=0
	for i:=0;i<5;i++{
		go addWithLock()
	}
	time.Sleep(time.Second)
	println("addWithLock",x)

}
//waitGroup优化HelloGoRoutine()
func ManyGoWait(){
	var wg sync.WaitGroup
	wg.Add(5)
	for i:=0;i<5;i++{
		go func(j int){
			defer wg.Done()
			Hello(j)
		}(i)
	}
	wg.Wait()
}