package Repository


import "sync"

//话题
type Topic struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Created_time int64 `json:"created_time"`

}


type TopicDao struct{

}
var (
	topicDao *TopicDao
	//一次性的初始化需要互斥量和boolean来记录初始化是否完成
	topicOnce sync.Once
)
func NewTopicDaoInstance() *TopicDao{
	//Do接收初始化函数作为参数
	//每次调用都会锁定Mutex并检测boolean变量 Do调用loadIcons并将boolean设为true
	topicOnce.Do(
		func(){
			topicDao=&TopicDao{}
		})
	return topicDao
}

func(*TopicDao) QueryTopicById(id int64) *Topic{
	//根据id获取相应内容
	return topicIndexMap[id]
}
