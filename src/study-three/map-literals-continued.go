/**
 * map的文法（续）
 * 如果顶级的类型只有类型名的话，可以在文法的元素中省略键名
 */
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.2345, -74.4590},
	"Google": {37.2391, -122.0723},
}

func main() {
	fmt.Println(m)
}