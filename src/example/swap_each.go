package main

import "fmt"

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	var str_1, str_2 = "abc", "xyz"
	str_1, str_2 = swap(str_1, str_2)
	fmt.Println("swap str_1 and str_2 ", str_1, str_2)
}
