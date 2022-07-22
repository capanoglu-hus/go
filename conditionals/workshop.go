package conditionals

import "fmt"

func Demo3() {
	// 3 adeet degişken int
	// ekrana en büyügü yazdır

	var a int = 25
	var b int = 82
	var c int = 104

	// var sayi1 , sayi2 , sayi3  int = 10 , 5 , 8

	if a < b {
		if b < c {
			fmt.Println(" c en büyüktür ")
		} else {
			fmt.Println("b en büyük ")

		}
	} else {
		fmt.Println("a en büyük")
	}
}
