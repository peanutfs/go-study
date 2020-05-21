/**
 * 图片
 * package image 定义了Image接口
 * package image
 * type Image interface {
 * 		ColorModel() color.Moder
 * 		Bounds() Rectangle
 * 		At(x, y int) color.Color
 * }
 * 注意：Bounds方法的Rectangle返回值实际上是一个image.Rectangle，其定义在image包中
 * color.Color和color.Model也是接口，但是通常因为直接使用预定义的实现image.RGBA和image.RGBAModel而不忽视了
 */
 package main

 import (
 	"fmt"
 	"image"
 )

 func main() {
 	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
 	fmt.Println(m.Bounds())
 	fmt.Println(m.At(0, 0).RGBA())
 }