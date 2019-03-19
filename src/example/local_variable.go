package main

import "fmt"

func main() {
	/* 声明局部变量 */
	var a, b, c int

	/* 初始化参数 */
	a = 10
	b = 20
	c = a + b

	fmt.Printf ("结果： a = %d and b = %d\n ret c = %d\n", a, b, c)
}
