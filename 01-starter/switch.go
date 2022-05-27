package main

import (
	"fmt"
	"time"
)

func main(){
	a := 2
	switch a {
	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
}
	t:=time.Now()
	switch  {
	case t.Hour()<12:
		fmt.Println("1")
	default:
		fmt.Println("2")
	}

	}
