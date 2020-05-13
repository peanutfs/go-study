/**
 * switch
 * 一个机构体就是一个字段的字和
 * 除非以fallthrough语句结束，否则分支会自动终止
 */

package main 

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s", os)
	}
}