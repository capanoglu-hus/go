package functions

// fonksiyonlarda aynı tipte olan ama
// parametre sayısı belli olmayan fonksiyonlar
//hesap makinesi toplama işleminde mesela 10 kere toplanır

func ToplaVariadic(sayilar ...int) int {
	// birden fazla parametre almış oldu bu formatta
	toplam := 0
	for i := 0; i < len(sayilar); i++ {
		toplam = toplam + sayilar[i]
	}
	return toplam
}
