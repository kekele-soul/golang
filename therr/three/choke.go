package main

import "fmt"

func main() {
	//阻塞的特点
	//等待输入直到用户敲了回车
	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scan(&age)
	fmt.Printf("你的年龄是：%d\n", age)
	/**

	 */
	//var num1 int
	//var num2 int
	//fmt.Println("请输入两个变量的值：")
	//fmt.Scan(&num1, &num2)
	//
	//var str string
	//
	//fmt.Println("请输入要执行的操作:(++、--)")
	//fmt.Scan(&str)
	//
	//switch str {
	//case "++":
	//	num1++
	//	fmt.Println("num1执行后的结果是：", num1)
	//
	//}
	//var ste string
	//fmt.Scan(&ste)
	//switch str {
	//case "++":
	//	num2++
	//	fmt.Println("num2执行后的结果是：", num2)
	//
	//}

}