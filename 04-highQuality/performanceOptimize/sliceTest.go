package main

import (
	"bytes"
	"strings"
)

func NoPreAlloc(size int){

	data:=make([]int,0)
	for k:=0;k<size;k++{
		data=append(data,k)
	}
}
func PreAlloc(size int){
	//初始化时指定大小
	data:=make([]int,0,size)
	for k:=0;k<size;k++{
		data=append(data,k)
	}
}
//操作slice
func GetLastBySlice(origin []int)[]int{
	return origin[len(origin)-2:]
}
func getLastByCopy(origin []int)[]int{
	result:=make([]int,2)
	copy(result,origin[len(origin)-2:])
	return result
}
//字符串拼接
func strBuilder(n int,str string)string{
	var builder strings.Builder
	for i:=0;i<n;i++{
		builder.WriteString(str)
	}
	return builder.String()
}
func ByteBuffer(n int,str string)string{
	buf:=new(bytes.Buffer)
	for i:=0;i<n;i++{
		buf.WriteString(str)
	}
	return buf.String()
}