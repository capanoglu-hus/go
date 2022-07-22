package functions

func DortIslem(sayi1 int, sayi2 int) (int, int, int, float64) {
	toplam := sayi1 + sayi2
	cikarim := sayi1 - sayi2
	carpim := sayi1 * sayi2
	bolum := float64(sayi1 / sayi2)

	return toplam, cikarim, carpim, bolum
	//birden fazla veriyi aynı anda döndürmek
	//diger dillerde encapsulation var
}
