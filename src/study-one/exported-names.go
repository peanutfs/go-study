/**
 * 在导入一个包之后，就可以用其导出的名称来调用它。
 * 在Go中，首字母大写的名称是被导出的。
 * Foo和FOO都是被导出的名称。名称foo是不会被导出的
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(math.pi) 执行会报错
	fmt.Println(math.Pi)
}