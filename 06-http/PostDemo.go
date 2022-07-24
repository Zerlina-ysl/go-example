package main

import (
	"context"
	"golang.org/x/tools/cmd/getgo/server"

)


func main(){
	h:=server.New()

	h.POST("/sis",func(c context.Context,ctx *app.RequestContext){
		ctx.Data(200,"text/plain;charset=utf-8",[]byte("OK"))
	})
	h.Spin()
}
