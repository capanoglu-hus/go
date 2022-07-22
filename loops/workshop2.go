package loops

import "fmt"

//bir kullanıcı bir sayı girsin
//asal olup olmadıgına bakacak

func Demo4() {

	sayi := 0
	fmt.Println("bir sayi gir ")
	fmt.Scanln(&sayi)

	asalMi := true

	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalMi = false
		}
	}

	if asalMi == true {
		fmt.Println("asal")
	} else {
		fmt.Println("asal degil")
	}
}
