/**
 * for是Go的while
 * 基于此可以省略分号：C的while在Go中叫做for
 */

package main

import "fmt"

func main() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}