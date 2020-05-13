/**
 * if语句除了没有()之外（强制不能使用），和C/Java中一样，{}必须
 */

 package main

 import (
 	"fmt"
 	"math"
 )

 func sqrt(x float64) string {
 	if x < 0 {
 		return sqrt(-x) + "i"
 	}
 	return fmt.Sprint(math.Sqrt(x))
 }

 func main() {
 	fmt.Println(sqrt(16), sqrt(-4))
 }