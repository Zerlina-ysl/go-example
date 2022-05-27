package Controller

import (
	"awesomeProject/03-enginerr/CommunityDemo/service"
	"strconv"
)

type PageData struct {

	Code int64 `json="code"`
	Msg string `json="msg"`
	Data interface{} `json:"data"`
}
//接收一个string类型的话题id，返回相应的页面数据（除了数据，还包含状态码和输出信息用于报告查询结果和错误等
func QueryPageInfo(topicIdStr string)*PageData {
	topicId,err:=strconv.ParseInt(topicIdStr,10,64)
	if err!=nil{
		return &PageData{
			Code:-1,
			Msg:err.Error(),
	}
	}
	pageInfo,err:=service.QueryPageInfo(topicId);
	if err!=nil{
		return &PageData{
			Code:-1,
			Msg:err.Error(),
		}
	}
	return &PageData{
		Code:0,
		Msg:"success",
		Data:pageInfo,
	}
}

