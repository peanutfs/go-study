/**
 * 接收者为指针的方法
 * 方法可以与命名类型或命名类型的指针关联
 * 刚刚看到的两个Abs方法。一个是在*Vertex指针类型上，而另一个在MyFloat值类型上
 * 有两个原因需要使用指针接收者，首先避免在每个方法调用中拷贝值（如果只值类型是大的结构体的话会更有效率）
 * 其次，方法可以修改接收者指向的值
 * 尝试修改Abs的定义，同时Scale方法使用Vertex代替*Vertex作为接收者
 * 当v是Vertex的时候Scale方法没有任何作用。
 * Scale修改v，当v是一个值（非指针），方法看到的是Vertex的副本，并且无法修改原始值
 * Abs的工作方式是一样的。只不过，仅仅读取v。所以读取的是原始值（通过指针）还是那个值的副本并没有关系。
 */
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}