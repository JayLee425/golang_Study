package main

import ("fmt")

func main()  {
	/* 定义局部变量 */
	var a, b, ret_max int
	a = 100
	b = 200
	/* 调用函数并返回最大值 */
	ret_max = max(a,b)
	fmt.Println("最大值max : %d { ", ret_max , " }" )
}

/** 返回两个数的最大值 **/
func max(num1, num2 int) int {
	var result_max_vale int
		if (num1 > num2) {
			result_max_vale = num1
		} else {
			result_max_vale = num2
		}
	return result_max_vale
}
