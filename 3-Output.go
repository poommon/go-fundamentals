package main

import (
	"fmt"
)

func main() {

	num := 20

	fmt.Println("Print number 20 of prinf() of function")
	fmt.Println("=====================================")
	fmt.Printf("%d %+d\n", num, num)
	fmt.Printf("%5d\n", num)
	fmt.Printf("%-5d%-5d\n", num, num)
	fmt.Printf("Number 70 is character : %c \n", num+50)
	fmt.Println("==============Scanf()=====================")
	fmt.Print("Please enter 3 number (format : num1,num2,num3) : ")
	var num1, num2, num3, sum int
	_, err := fmt.Scanf("%d,%d,%d", &num1, &num2, &num3)
	sum = num1 + num2 + num3
	if err == nil {
		fmt.Printf("Sumary of 3 numbers if %v\n", sum)
	} else {
		fmt.Printf("Error is \"%v\"\n", err)
	}
}
