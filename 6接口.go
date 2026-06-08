package main

import (
	"log"
)

// 定义一个接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义一个结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// 实现接口的方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 定义一个结构体
type Circle struct {
	Radius float64
}

// 实现接口的方法
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func main() {
	// 使用接口
	var s Shape
	s = Rectangle{Width: 10, Height: 5}
	log.Println("Area:", s.Area())
	log.Println("Perimeter:", s.Perimeter())

	s = Circle{Radius: 5}
	log.Println("Area:", s.Area())
	log.Println("Perimeter:", s.Perimeter())
}
