package main

import (
	"fmt"
)

// MyNum is a custom type based on int

func main() {
	fmt.Println("================ method=====================")
	n1 := MyNum(10)
	n2 := MyNum(20)

	// Calling the methods attached to the MyNum type
	fmt.Printf("%d + %d = %d\n", n1, n2, n1.Plus(n2))
	fmt.Printf("%d * %d = %d\n", n1, n2, n1.Multiply(n2))

	fmt.Println("================ Pointer receiver method=====================")
	m := MyMethod{myvar: 99}
	fmt.Println("Orignal , value is : ", m.myvar)
	m.ValueRec()
	fmt.Println("After call ValueRec(), value is : ", m.myvar)
	fmt.Println()
	fmt.Println("Orignal Pointer , value is : ", m.myvar)
	m.PointerRec()
	fmt.Println("After call PointerRec(), value is : ", m.myvar)

}

// ----------------------------------------------------
type MyNum int

func (num1 MyNum) Plus(num2 MyNum) MyNum {
	return num1 + num2
}

func (num1 MyNum) Multiply(num2 MyNum) MyNum {
	return num1 * num2
}

// ----------------------------------------------------
type MyMethod struct {
	myvar int
}

func (m1 MyMethod) ValueRec() {
	m1.myvar = m1.myvar * 10
}
func (m2 *MyMethod) PointerRec() {
	m2.myvar = m2.myvar * 500
}
