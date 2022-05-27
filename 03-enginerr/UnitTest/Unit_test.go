package UnitTest

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"testing"
)



func TestJudgePassLineTrue(t *testing.T){
	isPass:=JudgePassLine(70)
	assert.Equal(t,true,isPass )
}


func TestHelloTom(t *testing.T){
	output:= HelloTom()
	expectedOutput:="Tom"
	assert.Equal(t,expectedOutput,output)
	//if output!=expectedOutput{
	//	t.Errorf("expected %s* do not match actual %s",expectedOutput,output)
	//}
}

func TestProcessFirstLine(t *testing.T) {
	//firstLine := ProcessFirstLine()
	//打桩 defer卸载monkey
	//对ReadFirstLine打桩测试，不再依赖本地文件
	monkey.Patch(ReadFirstLine,func() string{
		return "line110"
	})
	defer monkey.Unpatch(ReadFirstLine)
	line:=ProcessFirstLine()
	assert.Equal(t, "line000",line)
}

//对select()进行基准测试
func BenchmarkSelect(b *testing.B) {
	InitServerIndex()
	//重置测试时间
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		//串行的压力测试
		Select()
	}
}
func BenchmarkSelectParallel(b *testing.B){
	InitServerIndex()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB){
		for pb.Next(){
			Select()
		}
	})
}
func BenchmarkFastSelect(b *testing.B) {
	FastSelect()
}