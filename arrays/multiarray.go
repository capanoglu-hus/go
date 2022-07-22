package arrays

import "fmt"

func Demo4() {
	var sayilar1 [2][3]int
	a := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			(sayilar1[i][j]) = a
			fmt.Print(sayilar1)
			fmt.Print(" ")

			a++
		}
		fmt.Println()

	}
	sayilar1[0][0] = 1
	sayilar1[0][1] = 3
	sayilar1[0][2] = 6
	sayilar1[1][0] = 9
	sayilar1[1][1] = 5
	sayilar1[1][2] = 2

	//fmt.Println(sayilar1[1][1])
}
