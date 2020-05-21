/**
 * goroutine
 * goroutine是由Go运行时环境管理的轻量级线程
 * go f(x, y, z)
 * 开启一个新的goroutine执行
 * f(x, y, z)
 * f, x, y和z是当前goroutine中定义的，但是在新的goroutine中运行f
 * goroutine在相同的地址空间中运行，因此访问共享内存必须进行同步，sync提供了这种可能，不过在Go中并不经常用到，因为有其他的办法
 */
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}