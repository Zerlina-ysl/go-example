package main

import (
	"fmt"
	"strings"
)

type point struct{
	x,y int
}
func main(){
	a:="hello"

	fmt.Println(strings.Contains(a,"ll"))
	fmt.Println(strings.Count(a,"l"))
	fmt.Println(strings.HasPrefix(a,"he"))

	fmt.Println(strings.HasSuffix(a,"lo"))

	fmt.Println(strings.Index(a,"ll"))

	fmt.Println(strings.Join([]string{"he","llo"},"-"))

	fmt.Println(strings.Replace(a,"e","E",-1))

	fmt.Println(strings.Repeat(a,2));

	fmt.Println(strings.Split("a-b-c","-"))

	fmt.Println(strings.ToLower(a))
	fmt.Println(strings.ToUpper(a))

	fmt.Println(len(a))
	b:="你好"

	fmt.Println(len(b))


	s:="hello"
	n:=123
	p:=point{1,2}
	fmt.Println(s,n)
	fmt.Println(p)

	fmt.Printf("s=%v\n",s)
	fmt.Printf("n=%v\n",n)
	fmt.Printf("p=%v\n",p)
	fmt.Printf("p=%+v\n",p)
	fmt.Printf("p=%#v\n",p)

	f:=3.1415926535
	fmt.Println(f)
	fmt.Printf("%.2f\n",f)






}
