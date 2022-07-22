package slice

import "fmt"

func Demo2() {
	sehirler := []string{"ankara", "istanbul", "izmir"}
	fmt.Println(sehirler)
	sehirlerKopya := make([]string, len(sehirler))
	// bellekte yer açmış oluyor make ile
	fmt.Println(sehirlerKopya)
	//sehirleri kopya kısmında nasıl aktaracaz
	//copy() fonksiyonuyla bir slice i başka bir slice a atıyor
	copy(sehirlerKopya, sehirler)
	//typelarına dikkat et
	fmt.Println(sehirlerKopya)
	sehirler = append(sehirler, "mersin")
	fmt.Println(sehirlerKopya)
	fmt.Println(sehirler)
	//filtrele işlemlerini
	fmt.Println(sehirler[0:3]) //3.indesten başlayıp 5 e kadar
	fmt.Println(sehirler[1:])
	// 1 den sonuna kadar aldı
	fmt.Println(sehirler[:3])
	//baştan 3 e kadar aldı

}
