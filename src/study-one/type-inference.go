/**
 * 类型推导
 * 在定义一个变量但不指定其类型时（使用没有类型的var或者:=语句），变量的类型由右值推导得出。
 * 当右值定义了类型时，新变量的类型与其相同
 * var i int
 * j := i // j也是一个int
 *
 * 但是当右边包含了未指名类型的数字常量时，新的变量就可能是int、float64或complex128，这取决于常量的精度
 * i :=42 // int
 * f := 3.142 // float64
 * g := 0.867 + 0.5i // complex128
 */

package main

import "fmt"

func main() {
	v := 23.3 
	fmt.Printf("v is of type %T\n", v)
}