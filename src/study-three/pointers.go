/**
 * 指针
 * Go具有指针。指针保存了变量的内存地址
 * 类型*T时指向类型T的指针，其零值时nil
 * var p *int
 *
 * &符号会生成一个指向其作用对象的指针
 * i := 42
 * p = &i
 *
 * *符号表示指针指向的底层的值
 * fmt.Println(*p) // 通过指针p读取i
 * *p = 21 // 通过指针p设置i
 * 这也就是通常所说的间接引用或非直接引用
 * 与C不同，Go没有指针运算
 */

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i // point to i
	fmt.Println(*p) // read i through the pointer

	*p = 21 // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j) // see the new value of j
}