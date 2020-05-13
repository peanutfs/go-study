/**
 * range
 * for循环的range格式可以对slice或者map进行迭代循环
 */
 package main

 import "fmt"

 var pow = []int{1, 2, 4, 8, 16, 32}

 func main() {
 	for i, v := range pow {
 		fmt.Printf("2**%d = %d\n", i, v)
 	}
 }