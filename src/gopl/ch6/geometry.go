package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

// 这是一个普通的函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 这是Point类型的方法，参数p是方法的接受者，直线距离
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println("函数调用:", Distance(p, q)) // "5" 函数调用
	fmt.Println("方法调用:", p.Distance(q))  // 方法调用
	fmt.Println("方法调用:", q.Distance(p))
}
