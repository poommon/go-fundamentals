package main

import (
	"fmt"
)

func main() {

	fmt.Println("================Struct=====================")
	a := [3]int{1, 2, 3}
	var p [3]*int
	//var i int
	p[0] = &a[0]
	p[1] = &a[1]
	p[2] = &a[2]

	for i := 0; i < len(p); i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *p[i])
	}

}
