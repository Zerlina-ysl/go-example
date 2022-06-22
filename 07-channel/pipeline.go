package main

import "fmt"

func main(){
	naturals1:=make(chan int)
	squares1:=make(chan int)
	go func(){
		for x:=0;;x++{
			naturals1<-x
		}
	}()
	go func(){
		for{
			x:=<-naturals1
			squares1<-x*x
		}
	}()
	fmt.Println(<-squares1)
}
