/**
 * 接口
 * 接口类型是由一组方法定义的集合
 * 接口类型的值可以存放实现这些方法的任何值
 */
 package main

 import (
 	"fmt"
 	"math"
 )

 type Abser interface {
 	Abs() float64
 }

 func main() {
 	var a Abser
 	f := MyFloat(-math.Sqrt2)
 	v := Vertex{3, 4}

 	// a MyFloat实现了Abser
 	a = f
 	// a *Vertex实现了Abser
 	a = &v
 	// 下面一行v是一个Vertex(而不是*Vertex)所以没有实现Abser
 	// a = v

 	fmt.Println(a.Abs())
 }

 type MyFloat float64

 func (f MyFloat) Abs() float64 {
 	if f < 0 {
 		return float64(-f)
 	}
 	return float64(f)
 }

 type Vertex struct {
 	X, Y float64
 }

 func (v *Vertex) Abs() float64 {
 	return math.Sqrt(v.X * v.X + v.Y * v.Y)
 }