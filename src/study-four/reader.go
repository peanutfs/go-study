/**
 * Readers
 * io包指定了io.Reader接口，它表示从数据流结尾读取
 * Go标准库包含了这个接口的许多实现，包括文件、网络连接、压缩、加密等等
 * io.Reader接口有一个Read方法
 * func (T) Read(b []byte) (n int, err error)
 * Read用数据填充指定的字节slice，并且返回填充的字节数和错误信息，在遇到数据流结尾时，返回io.EOF错误
 * 例子代码创建一个strings.Reader并且以每次8字节的速度读取它的输出
 */
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}