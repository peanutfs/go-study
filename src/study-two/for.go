/**
 * for
 * Go只有一种循环结构-for循环
 * 基本的for循环除了没有()之外（强制不能使用他们）
 * 看起来和C/Java中做的一样，而{}是必须的
 */

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}