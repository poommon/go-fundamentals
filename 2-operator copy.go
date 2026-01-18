package main

import (
	"fmt"
)

func main() {

	var x int = 100
	var y int = 50
	var z int = 100

	fmt.Println("===========operator =====")
	fmt.Println("x==z", x == z)
	fmt.Println("x!=y", x != y)
	fmt.Println("x>y", x > y)

	fmt.Println("=========== bitwise perator =====")
	var result int = x &^ y
	fmt.Printf("result = %08b\n", result)

}
