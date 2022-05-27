package service

import (
	"awesomeProject/03-enginerr/CommunityDemo/Repository"
	"errors"
	"sync"
)

type PageInfo struct {
	Topic *Repository.Topic
	PostList []*Repository.Post
}

//相等于查询的上下文context
type QueryPageInfoFlow struct{
	topicId int64
	pageInfo *PageInfo

	topic *Repository.Topic
	posts []*Repository.Post
}


func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId<=0{
		return errors.New("topic id must be larger than 0")
	}
	return nil
}

func (f *QueryPageInfoFlow) prepareInfo() error {
	//获取topic信息
	var wg sync.WaitGroup
	wg.Add(2)
	go func(){
		defer wg.Done()
		topic:=Repository.NewTopicDaoInstance().QueryTopicById(f.topicId)
		f.topic=topic
	}()
	//获取post列表
	go func(){
		defer wg.Done()
		post:=Repository.NewPostDaoInstance().QueryPostsByParentId(f.topicId)
		f.posts=post
	}()
	wg.Wait()
	return nil
}

func (f *QueryPageInfoFlow) packPageInfo() error {
	f.pageInfo=&PageInfo{Topic: f.topic,PostList: f.posts}
	return nil

}

func QueryPageInfo(topicId int64)(*PageInfo,error){
	//根据id给出页面信息，包括话题和回复内容
	return NewQueryPageInfoFlow(topicId).Do()
}

func NewQueryPageInfoFlow(topicId int64) *QueryPageInfoFlow{
	return &QueryPageInfoFlow{
		topicId: topicId,
	}
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	//检查id是否合法(>0
	if err:=f.checkParam();err!=nil{
		return nil,err
	}
	//通过id获取postList和topic
	if err:=f.prepareInfo();err!=nil{
		return nil,err
	}
	//将postList和topic赋值给PageInfoFlow
	if err:=f.packPageInfo();err!=nil{
		return nil,err
	}
	return f.pageInfo,nil
}

