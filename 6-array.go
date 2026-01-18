package main

import (
	"fmt"
)

func main() {

	fmt.Println("================Array=====================")
	var score = []float32{4, 3.5, 2, 3}
	var avg_cost float32 = 0.0
	var min_score float32 = 0.0

	for i := 0; i < len(score); i++ {
		fmt.Println(score[i])
		avg_cost = (avg_cost + score[i]) / 2

		//min_score = score[i]
		if i == 0 {
			min_score = score[i]
		}

		if min_score > score[i] {
			min_score = score[i]
		}
	}
	fmt.Println("avg cost :", avg_cost, " | min score : ", min_score)

	fmt.Println("================Array 2 Dimansion=====================")
	student := [3][2]string{{"MR.A", "3.5"}, {"MR.BB", "2"}, {"MR.CC", "4"}}

	fmt.Println("List of student : ")
	for x := 0; x < len(student); x++ {
		for y := 0; y < 2; y++ {
			fmt.Printf("[%d][%d] : %s\n", x, y, student[x][y])
		}
	}

}
