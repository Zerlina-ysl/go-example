package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main(){
	maxNum:=100
//用时间戳来初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	secretNumber:=rand.Intn(maxNum)
	fmt.Println("the secret number is",secretNumber)

	//读取用户输入
	fmt.Println("input your guess")
	reader := bufio.NewReader(os.Stdin)
	for {

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("an error ocurred while reading input,please try again", err)
		return
			continue
		}
		//去掉换行符
		input = strings.TrimSuffix(input, "\n")
		//转换成数字
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input.please enter an integer value")
		return
			//不知道两个continue存在的意义
			continue
		}
		fmt.Println("the guess is", guess)

		//判断逻辑实现
		if guess > secretNumber {
			fmt.Println("your guess is bigger than the secret number.please try again")
		} else if guess < secretNumber {
			fmt.Println("your guess is smaller than the secret number.please try again")

		} else {
			fmt.Println("Correct,you legend")
			break

		}
	}

}
