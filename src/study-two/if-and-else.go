/**
 * if和else
 * 在if的便捷语句定义的变量同样可以在任何对应的else块中使用
 */

package main

import (
	"fmt"
	"math"
)

func pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 8), pow(3, 2, 20))
}

