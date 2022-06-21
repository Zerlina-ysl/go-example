package main

import (
	"context"
	"golang.org/x/tools/cmd/getgo/server"
)

func main(){
	h:=server.New()


	h.Use(func(c context.Context,ctx *app.RequestContext){
		//print request
		logs.Infof("Received RawRequest:%s",ctx.Request.RawRequest())

		//next handler
		ctx.Next(c)
		//print response
		logs.Infof("Second RawResponse:%s",ctx.Response.RawResponse())
	})

	h.POST("/login",func(c context.Context,ctx *app.RequestContext){
		//some biz logic
		ctx.JSON(200,"OK")
	})
	h.POST("/logout",func(c context.Context,ctx *app.RequestContext){
		//some biz logic
		ctx.JSON(200,"OK")
	})
	h.spin()
}
//关于中间件的设计
//1.需要实现预处理和后处理，类似于函数的调用
func Middleware(some param){
	//pre-handle

	//调用下一个处理函数。
	Next()

	//after-handle

}
//路由上可以注册多middleware，也可以满足请求级别有效。只需将midlleware设计为业务和handler相同即可


//2. 如果用户不主动的调用下一个处理函数
func (ctx *RequestContext) Next(){
	ctx.index++
	for ctx.index<int8(len(ctx.handlers)){
		ctx.handlers[ctx.index]()
		//任何场景下index递增
		ctx.index++
	}
}

//3. 如果调用过程出现了异常，想要停止，如何处理
func (ctx *RequestContext) Abort(){
	//index设为最大值，直接跳出循环
	ctx.index=IndexMax
}