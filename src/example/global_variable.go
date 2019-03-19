package main
/* go语言全局变量 局部变量 名称相同时，优先选取局部变量 */
import "fmt"

/* 声明全局变量 */
var g int = 20

func main() {
	/* 声明局部变量 */
	var g int = 10

	fmt.Printf ("结果： g = %d\n",  g)
}
//todo 注意： result g = 10 （局部变量的值）