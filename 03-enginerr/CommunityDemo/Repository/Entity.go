package Repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func Init(filePath string) error{
	if error:=initTopicIndexMax(filePath);error!=nil{
		 fmt.Errorf("initTopicIndexMax error:%w",error)
		return error
	}
	if error:=initPostIndexMap(filePath);error!=nil{


		 fmt.Errorf("initPostIndexMap error:%w",error)
		 return error
	}
	return nil
}


var (
	//使用索引 数据行映射成map
	topicIndexMap map[int64]*Topic
	//Post的索引第一行表示的是 ParentId 也就是对应的Topic.Id
	postIndexMap map[int64][]*Post
)

func initTopicIndexMax(filePath string) error{
	//返回被打开的文件

	//fmt.Println("filePath:", filePath, "Topic.json")
	open,err:=os.Open(filePath+"Topic.json")
	if err!=nil{
		 fmt.Errorf("open file error:%w",err)
		return err
	}
	scanner:=bufio.NewScanner(open)
	topicTmpMap:=make(map[int64]*Topic)
	for scanner.Scan(){
		text:=scanner.Text()
		var topic Topic
		//把数据行转换为结构体存储
		if err:=json.Unmarshal([]byte(text),&topic);err!=nil{

			 fmt.Errorf("unmarshal failed:%w",err)
			 return err
		}
		topicTmpMap[topic.Id]=&topic
	}
	topicIndexMap=topicTmpMap
	return nil
}
func initPostIndexMap(filePath string) error{
	//返回被打开的文件
	open,err:=os.Open(filePath+"Post.json")
	if err!=nil{
		return err
	}
	scanner:=bufio.NewScanner(open)
	postTmpMap:=make(map[int64][]*Post)

	for scanner.Scan(){
		text:=scanner.Text()
		var  post Post
		//把数据行转换为结构体存储
		if err:=json.Unmarshal([]byte(text),&post);err!=nil{
			fmt.Errorf("unmarshal failed:%w",err)
			return err
		}
		posts,ok:=postTmpMap[post.Parent_id]
		if !ok{
			postTmpMap[post.Parent_id]=[]*Post{&post}
			continue
			}
		posts = append(posts,&post)
		postTmpMap[post.Parent_id]=posts

	}
	postIndexMap=postTmpMap
	return nil
}

