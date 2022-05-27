package main

import (
	"encoding/json"
	"fmt"
)

type userinfo struct{
	Name string
	Age int `json:"age"`
	Hobby []string
}
func main(){
	a := userinfo{Name:"wang",Age:18,Hobby: []string{"Golang","TypeScript"}}
	buf,err := json.Marshal(a)
	if err!=nil {
		panic(err)
	}
	fmt.Println(buf)
	fmt.Println(string(buf))

	buf,err=json.MarshalIndent(a,"","\t")
	if err!=nil {
		panic(err)
	}
	fmt.Println(string(buf))

	var b userinfo
	err = json.Unmarshal(buf,&b)
	if err!=nil {
		panic(err)
	}
	fmt.Printf("%#v\n",b)

}
