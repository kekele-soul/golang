package main

import "fmt"

/**
回顾go语言，第二章示例代码
 */
var num =30 //在函数外部定义 ，在程序的任何地方都能使用


/**变量的回顾
 变量就是一个代号,计算机当中的一小块内存空间，用来储存数据
格式：var 变量名 数据类型 = 变量的值
*/

func main() {
	var a int = 3	//定义一个整数变量,名字是a，值是3
	var b int = 4
	var c int
	//用公式计算面积
	c = (a * b) /2 //求c的值
	fmt.Println(c)
	fmt.Printf("c的类型是:%T\n",c)
	var d float32
	d = 5
	var e float32
	e = 2 * 3.14 * d
	fmt.Println("e的值是：",e)
	fmt.Printf("e的数据类型是:%T\n",e)
	//另一种定义方式
	age := 18
	fmt.Println("我的年龄是",age)
	fmt.Printf("age的数据类型是:%T\n",age)
	//go语言拥有类型推断的功能
	name := "可可"
	fmt.Printf("name的数据类型是:%T\n",name)


	//同时声明多个变量
	var h,i,j int
	h = 3
	i = 4
	j = 5
	fmt.Println(h,i,j)
	//对声明的变量修改
	h = 6
	i = 8
	j = 10
	fmt.Println(h,i,j)
	//另一种方式的声明多个变量
	var q,w,_ string = "liu","fei","yue"//_表示将该位置上的变量舍弃
	fmt.Println(q,w,)
	//第三种格式
	var na,ag,address  = "刘飞跃",18,"河南"
	fmt.Println(na,ag,address)
	fmt.Printf("na的数据类型：%T\n",name)
	fmt.Printf("na的内容是：%s,na的数据类型：%T\n",na,na)//格式化打印
	//第四种格式
	num1,num2 :=1,3
	fmt.Println(num1+num2)

	num1,num2 = 5,8
	fmt.Println(num1,num2)
	num1,num3 :=10,20
	fmt.Println(num1,num3)
	//变量的默认值
	//int默认值是0
	//string的默认值是""
	//float默认值是0




	//变量的集合
	var(
		name1 = "飞飞"
		age1 = 20
		address1 = "河南鹿邑"
		sex = "男"
	)
	fmt.Println(name1,age1,address1,sex)

	var age2 = 20//局部变量，只能在{}结尾前使用
	fmt.Println(age2)
	//调用chang函数
	chang()
}


//常量的复习
func chang(){
	//常量类似与变量，区别是常量不可改变
	//格式一；const 常量名字 数据类型 = 数值
	//常量的定义
	const PAI float32 = 3.14
	fmt.Println(PAI)
	fmt.Printf("a",PAI)
	//格式二
	const BAIDU = "www.baidu.com"
	fmt.Println(BAIDU)
	//规范：常量往往全部用大写来定义，用来和变量区分。

	//常量的集合
	const (
		MONDEY = "星期一"
		TUESDAY = "星期二"
		WENDENSDAY = "星期三"
	)

	//常量组
	const (
		 x int = 10
		 y
		 z
		 //y int = 10
		 //z int = 10
		 name1 string = "liufeiyue"
		 sex1 string = "男"
	)
	fmt.Println(x,y,z)

}