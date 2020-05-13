/**
 * 变量在定义是没有明确的初始化时会赋值为零值
 * 零值是：
 * 数值类型为0
 * 布尔类型为false
 * 字符串为""
 */

package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string 
	var c complex64
	fmt.Printf("%v %v %v %v %q\n", i, f, b, s, c)
}