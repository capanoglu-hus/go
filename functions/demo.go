package functions

import "fmt"

//sadece bir işlem yapmalı fonksiyon
//kendi kendini tekrar etme

func SelamVer(kullaniciAdi string) {
	fmt.Println("merhaba !!", kullaniciAdi)
}

func Topla() { //parametre göndererek yapacak olursak
	//buraya int türünde parametre yazman lazım
	sayi1 := 0
	fmt.Println("toplanacak ilk sayiyi gir")
	fmt.Scanln(&sayi1)
	sayi2 := 0
	fmt.Println("toplanacak ikinci sayiyi gir")
	fmt.Scanln(&sayi2)
	toplam := sayi1 + sayi2
	fmt.Println("girilen iki sayı toplamı : ", toplam)
}

func Topla2(sayi3 int, sayi4 int) int { //araya int yazmamızın sebebi dönüş tipini belirtir
	var toplam2 = sayi3 + sayi4
	fmt.Println("sonuc : ", toplam2)
	return toplam2
}
