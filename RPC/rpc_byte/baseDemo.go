package main

import "fmt"
/**
简单的本地调用
 */
func main(){
	var a=2
	var b=3
	result := calculate(a,b)
	fmt.Println(result)
	return
}
func calculate(a ,b int)int{
	return a*b
}

//struct Person{
////	Tag: Length					Value
//	1: required string 			userName,
//	2: optional i64   			favoriteNumber,
//	3: optional list<string>	interests
//}