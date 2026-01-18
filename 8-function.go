package main

import (
	"fmt"
	"strconv"
)

func main() {

	pprint("================function=====================")

	var mul_result int = 0
	mul_result = multiply(5, 20)
	pprint(strconv.Itoa(mul_result))

	pprint("================function cal=====================")
	a, b, c, d := cal(5, 1)
	pprint("add: " + strconv.Itoa(a))
	pprint("minus: " + strconv.Itoa(b))
	pprint("mul: " + strconv.Itoa(c))
	pprint("div: " + strconv.Itoa(d))

	pprint("================function Pointer=====================")
	var num1 int = 100
	var num2 int = 200
	funcPointer(&num1, &num2)
	pprint(" First Num: " + strconv.Itoa(num1))
	pprint(" Second Num: " + strconv.Itoa(num2))

	pprint("================function Array=====================")
	var myScore = [5]float32{3, 4, 3.5, 2, 2}
	var sum float32
	sum = funcSum(myScore)
	var str string = strconv.FormatFloat(float64(sum), 'f', 2, 32)
	pprint(" Sum value is : " + str)

	pprint("================function Closure=====================")
	printMsg("P'Poom")

	pprint("================User-Defined Function Types=====================")
	var sn sum_num = func(i1, i2 int) int {
		return i1 + i2
	}
	func_print(11, 5, sn)

}
func pprint(a string) {
	fmt.Println(a)
}
func multiply(a int, b int) int {
	return a * b
}
func cal(a, b int) (add int, minus int, mul int, div int) {
	add = a + b
	minus = a - b
	mul = a * b
	div = a / b
	return
}
func funcPointer(a *int, b *int) {
	*a = *a + 50
	*b = *b + 50
}
func funcSum(arr [5]float32) float32 {
	var sum float32
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
func printMsg(name string) {
	var hello = func() {
		pprint("Hello ! how are you : " + name)
	}
	hello()
}

type sum_num func(int, int) int

func func_print(num1 int, num2 int, s sum_num) {
	fmt.Printf("%d + %d = %d", num1, num2, s(num1, num2))
}
