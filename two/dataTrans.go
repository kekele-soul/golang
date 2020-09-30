package main

import (
	"fmt"
)
func main() {
	data()
}
//数据类型的转换
func data() {
	//数据类型转换，将数据强制转换
	//语法；
	//tyupe(value)
	//var num int
	//num = 100
	var num1 string
	num1 = "a"


	//sum :=int16(num) + num1//转换方式1
	//fmt.Println(sum)
	//sun := int(num1)
	fmt.Printf("num1的数据类型是：%T\n",num1)
	//grade := 13.14
	//reade := int8(grade)
	//fmt.Printf("数据类型是：%T\n",grade)
	//fmt.Printf("数据类型是：%T",reade)





	//关于float的小数点
	var pai = 3.1415926
	fmt.Printf("pai的数值是%f\n",pai)
	//保留两位小数
	fmt.Printf("pai的数值是%.2f\n",pai)
	//保留三位小数
	fmt.Printf("pai的数值是%.3f\n",pai)
	//保留一位小数
	fmt.Printf("pai的数值是%.1f\n",pai)

	//字符串
	name3 := "kekele"
	//查看字符串的长度
	fmt.Println("字符串的长度是：",len(name3))

	addr := "河南省周口市鹿邑县"
	//汉子和英语占用的空间不一样
	lenh :=len(addr)
	fmt.Println("地址的长度是：",lenh)
	//截取字符串
	name4 := "tiechuimeimei"
	//如何得到tiechui这个部分字符串
	ti := name4[:7]//截取前7个字符，从第一个开始可以省略不写
	fmt.Println(ti)
	//截取到末尾
	chui := name4[7:]//到末尾可以省略不写
	fmt.Println(chui)


}
