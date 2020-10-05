package main

import "fmt"

//多层for循环
func main() {
	//for i := 1; i <=3; i++{
	//	for j :=1; j <=4; j++ {
	//		fmt.Print("hello\t")
	//	}
	//	fmt.Println()
	//}


	//乘法表

	//for i :=1; i<=9; i++{
	//	for j := 1; j <=i; j++{
	//		fmt.Printf("%d*%d = %d\t",i,j,i*j)
	//	}
	//	fmt.Println()





	//打印菱形
	//for i := 1; i <=5; i++ {
	//
	//	for j := 1; j <= 5-i; j++ {
	//		fmt.Print(" ")
	//	}
	//	for sta := 1; sta <=2*i-1; sta++{
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}
	//for i :=1; i <=4; i++{
	//
	//	for num := 1; num <= i; num++ {
	//		fmt.Print(" ")
	//	}
	//	for sta :=1; sta <= 9-2*i; sta++{
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}
	
	
	//1000以内的水仙花数
	var i int
	for i = 0; i < 1000 ; i++  {
		 if i < 10{
		 	if i * i ==i{
		 		fmt.Println(i)
			}
		 	
		 }
		if i < 100 && i >10 {
			a := i % 10
			b := i / 10
			if a * a * a + b * b * b ==i{
				fmt.Println(i)
			}
		}
		if i >= 100 && i <=999{
			a := i / 100
			b := i / 10 %10
			c := i % 10
			if a * a * a + b * b *b + c * c *c ==i{
				fmt.Println(i)
			}
		}
	}
	
}
