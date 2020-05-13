/**
 * nil slice
 * slice的零值是nil
 * 一个nil的slice的长度和容量是0
 */
package main

import "fmt"

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}