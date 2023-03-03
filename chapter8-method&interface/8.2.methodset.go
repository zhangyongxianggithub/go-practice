//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-27 00:33
package main

import (
	"fmt"
	"math"
)

type Sphere struct {
	Radius float64
}

func (s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * s.Radius * s.Radius
}
func (s *Sphere) Volume() float64 {
	return float64(4) / float64(3) * math.Pi * s.Radius * s.Radius * s.Radius
}
func main2() {
	s := Sphere{
		Radius: 5,
	}
	fmt.Println(s.SurfaceArea())
	fmt.Println(s.Volume())
}
