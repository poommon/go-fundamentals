package main

import (
	"fmt"
)

func main() {
	fmt.Println("======hello interface ======")
	var t_shape, tra_shap Shape
	t_shape = Triangle{10, 9}
	fmt.Println("Traingle : area is ", t_shape.area())
	tra_shap = Trapezoid{2, 3, 5}
	fmt.Println("Trapezoid : area is ", tra_shap.area())
}

type Shape interface {
	area() float32
}
type Triangle struct {
	base, height float32
}
type Trapezoid struct {
	side1, side2, height float32
}

func (t1 Triangle) area() float32 {
	area := 0.5 * t1.base * t1.height
	return area
}
func (t2 Trapezoid) area() float32 {
	area := 0.5 * (t2.side1 + t2.side2) * t2.height
	return area
}
