/**
 * 错误
 * Go使用error值来表示错误状态
 * 与fmt.Stringer类似，error类型是一个内建接口
 * type error interface {
 * 		Error() string
 * }
 * (与fmt.Stringer类似，fmt包在输出是也会试图匹配error)
 * 通常函数会返回一个error值，调用的它的代码应当判断这个错误是否等于nil，来进行错误处理
 * i, err := strconv.Atoi("42")
 * if err != nil {
 * 		fmt.Printf("couldn't convert number: %v\n", err)
 * }
 * fmt.Println("Converted integer:", i)
 *
 * error为nil时表示成功，非nil的error表示错误
 */
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(), "it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}