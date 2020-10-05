package main

import "fmt"

/**
流程语句的复习
 */
func main() {
 	//if的运用

	var num int
	num = 59
	if num <= 60 {
		fmt.Printf("成绩不合格是：%d分\n",num)

	}
	fmt.Println("程序结束")

	fmt.Println("程序开始执行")
	age := 22
	if age >= 22 {
		fmt.Println("可以结婚了")
	}
	fmt.Println("执行结束")

	//判断是否能够结婚
	age1 := 19
	if age1 >= 22 {
		fmt.Println("可以结婚了")
	}else{
		fmt.Println("对不起还不能结婚")
	}
	//if语句的嵌套
	var age3 int
	fmt.Println("请输入你的年龄")
	fmt.Scan(&age3)
	var sex int
	fmt.Println("请输入性别，男的输入1女的输入2")
	fmt.Scan(&sex)
	if sex ==1  {
		if age3 >= 22 {
			fmt.Println("小哥哥可以结婚了")
		}else {
			fmt.Println("还不能结婚")
		}
	}else if sex  == 2{
		if age3 >= 20{
			fmt.Println("小姐姐可以结婚了")
		}else {
			fmt.Println("小姐姐再等等")
		}
	}



}
