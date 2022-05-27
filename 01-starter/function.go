package main

import "fmt"

func add(a int,b int)int {
	return a+b
}
func add2(a,b int) int{
	return a+b
}
func add3(n int){
	n+=2
}
func add3ptr(n *int){
	*n+=2
}
func exists(m map[string]string,k string)(v string,ok bool){
	v,ok=m[k]
	return v,ok
}
func main(){
	res:=add(1,2)
	fmt.Println(res)

	v,ok:=exists(map[string]string{"a":"A"},"a")
	fmt.Println(v,ok)

	n:=5
	add3(n)
	fmt.Println(n)
	add3ptr(&n)
	fmt.Println(n)

}
