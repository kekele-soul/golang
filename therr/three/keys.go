package main

import "fmt"

//键盘输入
func main() {


	//age := 19
	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scan(&age)//读取键盘上的输入
	fmt.Println(&age)//&age是获取age的内存地址
	if age >= 22 {
		fmt.Println("可以结婚")
	}else {
		fmt.Println("不能结婚再等等")
	}

	var a int
	var b int
	fmt.Println("请输入两个数，给a和b进行赋值")
	fmt.Scan(&a,&b)
	fmt.Printf("a的值是%d,b的值是%d\n",a,b)

	//用户名登录

	var uesrname string
	var password string
	fmt.Println("请输入你的用户名")
	fmt.Scan(&uesrname)
	fmt.Println("请输入你的密码")
	fmt.Scan(&password)
	if uesrname == "1823718761"{
		if password == "kekele" {
			fmt.Println("登陆成功")
		}
	}else {
		fmt.Println("登陆失败")
	}
	if c := 5; c > 0 {
		fmt.Printf("b的值是%d",c)
	}
}