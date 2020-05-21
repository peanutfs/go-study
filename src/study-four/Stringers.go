/**
 * Stringers
 * 一个普遍存在的接口是fmt包中定义的Stringer
 * type Stringer interface {
 * 		String() string
 * }
 *
 * Stringer 是一个可以用字符串描述自己的类型。fmt包（还有许多其他包）使用这个来进行输出
 */
package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}