/**
 * range和close
 * 发送者可以close一个channel来表示再没有值会被发送了，
 * 接收者可以通过赋值语句的第二参数来测试channel是否被关闭
 * 当没有值可以接收并且channel已经被关闭，那么经过v, ok := <-ch
 * 之后ok会被设置为false
 * 循环for i := range c 会不断的从channel接收值，直到它被关闭
 * 注意：只有发送者才能关闭channel，而不是接收者，向一个已经关闭的channel发送数据会引起panic，
 * 还要注意：channel与文件不同，通常情况下无需关闭他们，只有在需要告诉接收者没有更多数据的时候才有必要惊喜关闭，
 * 例如中断一个range
 */
package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}