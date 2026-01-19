package main

import (
	"fmt"
	"mycalculator/calculator/mycal"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("Plus: ", mycal.Plus(10, 5))
	fmt.Println("Minus: ", mycal.Minus(10, 5))
	fmt.Println("Multiply: ", mycal.Multiply(10, 5))
	fmt.Println("Divide: ", mycal.Divide(10, 5))
	fmt.Println("Modulus: ", mycal.Modulus(10, 5))

}
