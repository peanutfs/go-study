/**
 * for(续)
 * 跟C/Java中一样，可以让前置、后置语句为空
 */

package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 100; {
		sum += sum
	}

	fmt.Println(sum)
}