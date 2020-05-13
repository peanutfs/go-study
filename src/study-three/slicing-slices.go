/**
 * 对slice切片
 * slice可以重新切片，创建一个新的slice值指向相同的数组
 * 表达式
 * s[start:end]
 * 表示从start到end-1的slice元素，含两端。因此
 * s[start:start]是空的，而
 * s[start:start+1]有一个元素
 */

package main

import "fmt"

func main() {
	p := []int{2, 3, 4, 5, 6, 7}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4]==", p[1:4])

	// 省略下标表示从0开始
	fmt.Println("p[:3] ==", p[:3])

	// 省略上标代表到len(s)结束
	fmt.Println("p[4:] ==", p[4:])
}