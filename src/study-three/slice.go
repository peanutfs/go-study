/**
 * slice
 * 一个slice会指向一个序列的值，并且包含来了长度信息
 * []T是一个元素类型为T的slice
 */

package main

import "fmt"

func main() {
	p := []int{2, 3, 4, 5, 6}
	fmt.Println("p == ", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}