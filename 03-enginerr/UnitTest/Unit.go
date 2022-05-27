package UnitTest

import (
	"bufio"
	"github.com/NebulousLabs/fastrand"
	"math/rand"
	"os"
	"strings"
)

func JudgePassLine(score int16)bool{
	if score>=60{
		return true
	}
	return false
}

func HelloTom() string{
	//return "Tom"
	return "jerry"
}


func ReadFirstLine() string{
	open,err:= os.Open("log")
	defer open.Close()
	if err!=nil{
		return ""
	}
	scanner :=bufio.NewScanner(open)
	for scanner.Scan(){
		//对每行进行遍历
		return scanner.Text()
	}
	return ""
}
func ProcessFirstLine() string{
	line:=ReadFirstLine()
	//将行中包含11的数据进行替换为00
	destLine :=strings.ReplaceAll(line,"11","00")
	return destLine
}

//reflect.Value可以装载任意类型的指，ValueOf接受任意的interface{}，并返回一个装载着动态值的reflect.Value
//func Patch(target,replcement interface{}) *PatchGuard{
//	t:=reflect.ValueOf(target)
//	r:=reflect.ValueOf(replcement)
//	patchValue(t,r)
//	return &PatchGuard(t,r)
//}
//func Unpatch(target interface{})bool{
//	return unpatchValue(reflect.ValueOf(target))
//}

var ServerIndex [10]int

func InitServerIndex(){
	for i:=0;i<10;i++{
		ServerIndex[i]=i+100
	}

}
func Select() int{
	//随机选择执行服务器
	return ServerIndex[rand.Intn(10)]
}

func FastSelect() int{
	return ServerIndex[fastrand.Intn(10)]
}