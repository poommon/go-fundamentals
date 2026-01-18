package main

import (
	"fmt"
)

func main() {

	var month int

	fmt.Print("Please enter month (Ex:Jan=01 , Dec=12) :")
	fmt.Scan(&month)

	fmt.Println("================Swith Case=====================")
	switch month {
	case 1, 2, 3:
		fmt.Println("This is Q1 (First Quarter) ")
	case 4, 5, 6:
		fmt.Println("This is Q2 (Second Quarter) ")
	case 7, 8, 9:
		fmt.Println("This is Q3 (Third Quarter) ")
	case 10, 11, 12:
		fmt.Println("This is Q4 (Fourth Quarter) ")
	default:
		fmt.Println("%s not found ", month)
	}
	fmt.Println("================ For Loop=====================")
	for i := month; i <= 12; i++ {
		fmt.Printf("10*%d \t = \t %d\n ", i, 10*i)
	}
}
