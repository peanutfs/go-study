/**
 * channel
 * channel是有类型的管道，可以用channel操作符<-对其发送或者接收值
 * ch <- v // 将v送入channel ch
 * v := <-ch // 从ch接收，并且赋值给v
 * 箭头就是数据流的方向
 * 和map与slice一样，channel使用前必须创建
 * ch := make(chan int)
 * 默认情况下，在另一端准备好之前，发送和接收都会阻塞，这使得goroutine可以在没有明确的锁或静态变量的情况下进行同步	
 */
package main

import "fmt"

func sum(a []int,c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	// 将和送入c
	c <- sum 
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	// 从c中获取
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}