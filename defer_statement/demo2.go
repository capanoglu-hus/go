package deferstatement

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc()
	if sayi%2 == 0 {
		fmt.Println("çift sayi çalıstı")
		return "Çift sayıdır" //sonuıca deger atandı
	}

	if sayi%2 != 0 {
		return "tek sayıdır"
	}

	return "belli degil"
}

func Test() {
	sonuc := Demo2(9)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("bitti")
}
