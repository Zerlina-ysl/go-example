package main

import (
	"fmt"
)

type user1 struct{
	name string
	password string
}

func main(){
	a := user{name:"li",password:"1024"}
	b := user{"li","1024"}
	c :=user{name:"li"}
	c.password="1024"
	var d user
	d.name="wang"
	d.password="1024"

	fmt.Println(a,b,c,d)

	fmt.Println(checkPassword(a,"haha"))
	fmt.Println(checkPassword2(&a,"haha"))

	a.resetPassword("2048")
	fmt.Println(a.checkPassword("2048"))

}
func checkPassword(u user,password string) bool{
	return u.password==password
}
func checkPassword2(u *user,password string) bool {
	return u.password==password
}

func (u *user)resetPassword(password string){
	u.password=password
}
func (u user)checkPassword(password string)bool{
	return u.password==password
}