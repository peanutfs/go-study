/**
 * 结构体文法
 * 结构体文法表示通过结构体字段的值作为列表来新分配一个结构体
 * 使用Name: 语法可以仅列出部分字段（字段名的顺序无关）
 * 特殊的前缀& 返回一个指向结构体的指针
 */

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v = Vertex{1, 2} // 类型为Vertex
	x = Vertex{X: 1} // Y: 0被忽略
	z = Vertex{} // X:0和Y:0被忽略
	p = &Vertex{1, 2} // 类型为*Vertex
)

func main() {
	fmt.Println(v, x, z, p)
}
