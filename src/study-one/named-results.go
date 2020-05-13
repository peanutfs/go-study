/**
 * go的返回值可以被命名，并且想变量那样使用
 * 返回值的名称应当具有一定的意义，可以作为文档使用
 * 没有参数的return语句返回结果的当前值，也就是直接返回
 * 直接返回语句仅应用在像下面这样的短函数中，在长函数中他们会影响代码可读性
 */

package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}