package main

import "fmt"

func main() {
	oper()
}
//运算符学习
func oper()  {

	//数学中的运算：加法、减法、乘法、除法（+、-、*、/）

		//算术运算符：+、-、*、/
		//++ 在原来的基础上加一
		//-- 在原来的基础上减一
	a := 3
	b := 5
	sum := a+b
	fmt.Println(sum)
	sub := a - b
	fmt.Println(sub)
	su := a * b
	fmt.Println(su)
	suq := a / b
	fmt.Println(suq)
	//取模、取余数
	mod := b % a
	fmt.Println(mod)

	//自增和自减
	age := 20
	fmt.Println("现在的年龄：",age)
	age++
	fmt.Println("过了一年的年龄：",age)
	age--
	age--
	fmt.Println("穿越回去了两年的年龄：：",age)

	//关系运算符 > < = >= <= == === !=
	//运算的结果只能是 true 或者 false
	age1 := 18
	age2 := 19
	resule :=age1 > age2//结果为false
	fmt.Println(resule)

	resule1 := age1 < age2//j结果为true
	fmt.Println(resule1)

	//==判断两个值是否相等
	num1 := 100
	num2 := 100
	resule2 := num1 == num2//true
	fmt.Println(resule2)
	//!=： 不等于
	resule3 := num1 != num2//fllse
	fmt.Println(resule3)



	/**逻辑运算符:与、或、非
		与：&& （并且，同时成立）
		规则:操作的两边都是bool类型
		规律：全真则真、一假则假

		或：|| （或者，成立一个即可）
	 */
	b1 := false
	b2 := false
	res := b1 && b2 //false && false
	fmt.Println(res)
	b3 := true
	b4 := false
	res1 := b3 && b4 //true && false
	fmt.Println(res1)//结果是false

	num3 := 8
	num4 := 10
	//false && true 结果是false
	res2 := num3>=num4 && num3 <num4
	fmt.Println(res2)
		//true && false && false 结果是 false
	res3 := num3 != num4 && num3 > num4 && num3 >= num4
	fmt.Println(res3)


	//逻辑或
	c := 3
	d := 4
	//一真则真、全假则假
	resu := c < d || c == d  || c > d//结果为true
	fmt.Println(resu)

	//逻辑非、！=不等于
	resu1 := !(c < d)
	fmt.Println(resu1)




	//位运算符： 计算机中的二进制表示只有0和1 ：int8、 int16；位数
		//十进制 --》转换成二进制
	//按位与
	nub := 2
	nub1 := 7
	re1 := nub & nub1
	fmt.Println(re1)
	// 按位或
	nb := 4
	nb1 := 7
	re2 := nb | nb1
	fmt.Println(re2)
	//按位异或
	nb2:= 6
	nb3 := 9
	re3 := nb2  ^ nb3
	fmt.Println(re3)


	//按位向左移动
	ke := 3 //二进制是11
	re4 := ke <<2
	fmt.Println(re4)

	//按位向右移动
	ke1 := 7 //
	res4 := ke1 >> 2//向右移动两位
	fmt.Println(res4)








	//赋值运算符
	var k int //初始值0
	k = 3		//赋值
	fmt.Println(k)//3
	//k+一个值
	k += 2
	fmt.Println(k)
	//k-一个值
	k -= 2
	fmt.Println(k)
	//k*一个值
	k *= 2
	fmt.Println(k)
	//k/一个值
	k /= 2
	fmt.Println(k)
	k %= 3
	//按位与
	k &= 3
	//等等。。

	//交换两个值
	f := 3
	f1 := 5
	fmt.Println(f,f1)
	f2 := f
	f = f1
	f1 = f2
	fmt.Println(f,f1)
	//第二种方法
	age3,age4 := 15,18
	fmt.Println(age3,age4)
	age3,age4 = age4 ,age3//多变量交换
	fmt.Println(age3,age4)

}