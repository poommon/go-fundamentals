package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("================String Join=====================")
	var mystr = []string{"red", "green", "blue"}
	var joint = ","
	var str_result1 = strings.Join(mystr, joint)
	fmt.Println("result joint : ", str_result1)

	fmt.Println("================String Split=====================")
	var mystr2 = "red,green,yello"
	var str_result2 = strings.Split(mystr2, joint)
	fmt.Println("result split string : %s\n", str_result2)

	fmt.Println("================Replace=====================")
	var str_result3 = strings.Replace(mystr2, "green", "==Null==", 1)
	var str_result4 = strings.ReplaceAll(mystr2, "green", "==all==")
	fmt.Println("Result Replace : %s\n", str_result3)
	fmt.Println("Result Replace all : %s\n", str_result4)

}
