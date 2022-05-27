package main

import (
	"awesomeProject/03-enginerr/CommunityDemo/Controller"
	"awesomeProject/03-enginerr/CommunityDemo/Repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main(){
	if err:=Repository.Init("./data/");err!=nil{

		fmt.Errorf("error filepath:%w",err)
		os.Exit(-1)
	}
	//先创建树的根结点
	//生成Engine
	r:=gin.Default()
	//每种方法对应一棵独立的压缩检索树，树之间不共享数据
	r.GET("/comment/page/get:id",func(c *gin.Context){
		topicId:=c.Param("id")
		data:=Controller.QueryPageInfo(topicId)
		//输出json数据
		c.JSON(200,data)
		//c.JSON(200,
		//	Controller.QueryPageInfo(
		//		c.Param("id")))
	})
	err:=r.Run()
	if err!=nil{

		fmt.Errorf("error run:%w",err)
		return
	}
}

func Init(filePath string)error{
	if err:=Repository.Init(filePath);err!=nil{
		return err
	}
	return nil
}