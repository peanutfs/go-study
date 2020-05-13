/**
 * 类型转换
 * 表达式T(v)将值v转换为类型T
 * 一些关于数值的转换
 * var i int = 42
 * var f float64 = float64(i)
 * var u uint = uint(f)
 * 或者，更加简单的形式
 * i := 42
 * f := float64(i)
 * u := uint(f)
 * 与C不同的是Go的在不同类型之前的项目赋值时需要显式转换	
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x * x + y * y))
	var z int = int(f)
	// var z int = f 必须显式转换，否则会报错
	fmt.Println(x, y, z)
}