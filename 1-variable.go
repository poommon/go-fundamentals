package main

import (
	"fmt"
)

func main() {
	var str string = "Test"
	var num int = 20
	num2 := -100.55
	var status bool = false
	var (
		prd      = "P001"
		prd_name = "T-Shrit"
		price    = 250
	)

	//str := "Test"
	//num := 20
	//status := true
	emoji := '🐹'
	fmt.Printf("===========Variable %c =================\n", emoji)
	fmt.Println("Test variable : ", str, num, num2, status)
	fmt.Println("Product : ", prd, " | ", prd_name, " | ", price)

	fmt.Println("\n===========PI cicum=================")
	const pi = 3.14
	var cal_cicum = 2 * pi * 2
	fmt.Println("circum = ", cal_cicum)

	fmt.Println("\n============= Complex Number ==============")
	var aa float32 = 5
	var b float32 = 10
	x := 20 + 30i
	var y complex64 = complex(1, 2)
	z := complex(aa, b)
	addComplex := y + z
	fmt.Println("Complex number : ", x, " | ", y, " | ", z, real(addComplex), imag(addComplex))

}
