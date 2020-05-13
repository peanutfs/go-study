/**
 * if便捷语句
 * 和for一样，if语句可以在条件之前执行一个简单的语句
 * 由这个语句定义的变量的作用域仅在if范围之内
 */

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 8), pow(3, 2, 20))
}