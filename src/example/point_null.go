package main

import "fmt"

var flag int

func main() {
	var ptr *int
	if (ptr != nil) {
		flag = 1 /* ptr 不是空指针 */
	}
	if (ptr == nil) {
		flag = 0 /* ptr 是空指针 */
	}
	fmt.Printf("ptr 的值为 : %x\n, flag : %x\n", ptr, flag)
}
