package for_range

import "fmt"

func Demo2() {
	//1-10 arasında sayılardan tek olanları topla
	toplam := 0
	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, sayi := range sayilar {
		if sayi%2 == 1 {

			toplam = toplam + sayi
			fmt.Println("toplam", toplam)
		}

	}
}
