/**
 * 函数
 * 函数可以没有参数或接受多个参数
 * 在这个例子中，add接受两个int类型参数
 * 注意类型在变量名之后
 */

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42,13))
}