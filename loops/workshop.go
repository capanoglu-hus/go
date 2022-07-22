package loops

import "fmt"

// 1-100 arası sayı tahmin oyunu
func Demo3() {
	aklimda := 78
	tahmin := 0

	fmt.Println("bir sayi tahmin et ")
	fmt.Scanln(&tahmin)
	fmt.Println(tahmin)

	a := 0
	b := 0
	for aklimda != tahmin {
		if tahmin < aklimda {

			fmt.Println("DAHA büyük bir sayi gir  ")
			fmt.Scanln(&tahmin)
			a++
		}
		if tahmin > aklimda {

			fmt.Println("DAHA küçük bir sayi gir  ")
			fmt.Scanln(&tahmin)
			b++
		}
	}

	if tahmin == aklimda {
		fmt.Println("kazandın")
		toplam := a + b + 1
		fmt.Println(fmt.Sprintf("%v", toplam), ". tahminde  bildin")
		if toplam <= 3 {
			fmt.Println("süper")
		} else if toplam <= 6 {
			fmt.Println("güzel")
		} else {
			fmt.Println("fena degil")

		}

	}

}
