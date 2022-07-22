package variables

import "fmt" //formattan geliyor
//dry kendinii tekrar etme

func Demo1() {
	var metin string = "merhaba !! "
	fmt.Println(metin)
	fmt.Print(metin)

	var kdv int = 23
	fmt.Println(kdv)
	fmt.Println(100 + (100 * 10 / 100))

	var kdv2 float32 = 0.1
	fmt.Println(kdv2)
	fmt.Println(100 + 100*kdv2)

	kdv3 := 55.2 // : hem tanımlayıp hem atama yapıyoruz

	fmt.Println(kdv3)
	fmt.Printf("veri tipi : %T ", kdv3)

	var durum bool

	var metin1 string = "engin"
	var metin2 string = "ahmet"

	durum = metin1 == metin2
	fmt.Println(durum)
	fmt.Println(!durum)
}
