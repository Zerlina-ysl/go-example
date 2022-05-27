package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main(){
	rpc.RegisterName("HelloService",new(HelloService))
	//架设在jsonrpc路径
	//基于http.ResponseWriter和http.Request类型的参数构造io.ReadWriteCloser类型的conn通道
	http.HandleFunc("/jsonrpc",func(w http.ResponseWriter,r *http.Request){
		var conn io.ReadWriteCloser= struct {
			io.Writer
			io.ReadCloser
		}{
			//基于构建的conn构建针对服务器端的JSON编码解码器
			ReadCloser: r.Body,
			Writer: w,
		}
		//通过rpc.ServeReuqest()为每次请求处理一次RPC方法调用
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234",nil)
}