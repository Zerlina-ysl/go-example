package Repository

import "sync"
//定义实体结构
//回帖 Post.json.ParentId == Topic.Id Post的id是独立于其他id
type Post struct {
	//	"id":1,
	//	"title":"青训营来啦！",
	//	"content":"怎么都爱搞性别，是不是太恶臭了",
	//	"create_time":1650437625

	Id int64 `json:"id"`
	Content string `json:"content"`
	Created_time int64 `json:"created_time"`
	Parent_id int64 `json:"parent_id"`

}
type PostDao struct {

}
var(
	postDao *PostDao
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDao{
	postOnce.Do(
		func(){
			postDao=&PostDao{}
		})
	return postDao
}
func(*PostDao) QueryPostsByParentId(parentId int64)[]*Post{
	return postIndexMap[parentId]
}

