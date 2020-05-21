/**
 * 缓冲channel
 * channel可以是_带缓冲的_。为make提供第二参数作为缓冲长度来初始化1一个缓冲channel
 * ch := make(chan int, 100)
 * 向缓冲channel发送数据的时候，只有在缓冲区满的时候才会阻塞。当缓冲区清空的时候接受阻塞
 */
package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}