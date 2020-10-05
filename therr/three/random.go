package main

import (
	"fmt"
	"math/rand"
	"time"
)

//随机数和时间包的使用
func main() {

//num := rand.Intn(5 )//生成随机数的范围和类型
//	fmt.Println(num)
//	nub := rand.Intn(5 )//生成随机数的范围和类型
//	fmt.Println(nub)
//	num1 := rand.Intn(10 )//生成随机数的范围和类型
//	fmt.Println(num1)

	//seed()种子值

	//获取当前时间戳
	//time1 := time.Now()
	//time1.Unix()
	//time1.UnixNano()
	//rand.Seed(time1.UnixNano())
	//nem := rand.Intn(10)
	//fmt.Println(nem)
	rand.Seed(time.Now().UnixNano())
	nem := rand.Intn(10)
	fmt.Println(nem)
}
