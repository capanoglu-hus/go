package arrays

import "fmt"

func Demo3() {
	sayilar := [5]int{20, 30, 50, 10, 2}
	fmt.Println(sayilar)

	enBuyuk := sayilar[0] //ilk elemanı en büyük sayı olarak kabul edilir

	//length = uzunluk
	for i := 0; i < len(sayilar); i++ {
		if sayilar[i] > enBuyuk {
			enBuyuk = sayilar[i]
		}

	}
	fmt.Println("en büyük : ", enBuyuk)
}
