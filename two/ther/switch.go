package main

import "fmt"

//switch分支语句
func main() {
	//语法结构：
		//switch 变量{
			//case 值1：
			//case 值2：
			//...
	//fallthrough穿透 执行下一个
	//break，用于强制结束某个case的执行

	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scan(&age)
	switch age {
	case 20:
		fmt.Println("如果你是女生那你可以结婚了。")
	case 22:
		fmt.Println("不管你是男的还是女的都可以结婚了")
	default:
		fmt.Printf("你的年龄%d",age)
	}

	//一年四季
	var num int
	fmt.Println("请输入1-4的一个数")
	fmt.Scan(&num)
	switch num {
	case 1:
		fmt.Println("春天")
	case 2:
		fmt.Println("秋天")
	case 3:
		fmt.Println("夏天")
	case 4:
		fmt.Println("冬天")
	/*default:
		fmt.Println("程序运行结束")*/
	}
	fmt.Println("请输入两个数：")
	var num1 int
	var num2 int
	fmt.Scan(&num1,&num2)
	fmt.Println("请输入要执行的操作（+、-、*、/")

	var str string
	fmt.Scan(&str)
	switch str {
	case "+":
		fmt.Println("两数之和是：",num1+num2)

	}
}