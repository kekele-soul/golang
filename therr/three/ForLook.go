package main

import "fmt"

//for循环的复习
//break终止循环、无论条件是否满足
//continue结束当前本次循环、继续下次循环
func main() {
	sum := 0
	for i := 0 ;i <=100; i++{
		sum += i
	}
	fmt.Println("1-100的和是：",sum)


	for j := 1; j<= 20; j++{
		if j % 3 == 0 {
			fmt.Println(j)
		}
	}
	//5的阶乘
	res := 1
	for i := 1; i <= 5 ;i++  {
		res *= i
	}
	fmt.Println("5的阶乘是：",res)

	//表达式可以省略不写
	var a int
	for ; a < 10;{
		fmt.Println(a)
		a++
	}
	//不写条件默认是true，会一直执行，直到死机。
	//for  {
	//	fmt.Println("hello")
	//}





}
