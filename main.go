package main

import (
	"fmt"
	"golesson/project"
)

//camelCase ilk kelime küçük sonraları büyük
//PascalCase ilk harfi büyük

func main() {
	//fmt.Println("hey")
	// variables.Demo1()
	// conditionals.Demo3()
	//loops.Demo5()
	// slice.Demo2()
	// functions.SelamVer("hç")
	// functions.Topla()
	// sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(10, 5)
	// istemedigin zaman _ koyarsan onu alamış olursun
	// fmt.Println("toplam : ", sonuc1)
	// fmt.Println("çıkarma : ", sonuc2)
	// fmt.Println("çarpma : ", sonuc3)
	// fmt.Println("bölüm : ", sonuc4)

	// fmt.Println(functions.ToplaVariadic(4, 5, 33, 7, 6, 4))

	// sayis := []int{4, 5, 7, 0, 11}
	// fmt.Println(functions.ToplaVariadic(sayis...))
	// maps.Demo()
	// for_range.Demo3()
	// sayi := 20
	// pointer.Demo(&sayi)                  // yukarıdaki degeri direkt demoya aktarıyor ve işlem yapıyor
	// fmt.Println("maindeki sayı =", sayi) // direkt degeri gönderiyor
	// & pointere gödereceksen bu işaret olmalı

	// sayilar := []int{1, 2, 3}
	// pointer.Demo2(sayilar)
	// fmt.Println("maindeki sayı =", sayilar[0])
	// arrayler bellek yerine göre çalışır ( adres )
	// structs.Demo2()
	// CiftSayiCn := make(chan int) // sayı taşıyaan bir channel gibi
	// TekSayiCn := make(chan int)
	// go channels.CiftSayilar(CiftSayiCn)
	// go channels.TekSayilar(TekSayiCn)

	// ciftSayiToplam, tekSayiToplam := <-CiftSayiCn, <-TekSayiCn
	// // başka bir iş olarak onları açıyor
	// // ama kendi bittigi için digerlerini umursamıyor

	// time.Sleep(5 * time.Second)
	// carpim := ciftSayiToplam * tekSayiToplam
	// fmt.Println("Çarpım :", carpim)
	// interfaces.Demo3()
	// fmt.Println(errorhandling.TahminEt2(10))
	// stringfunctions.Demo2()
	// restful.Demo2()
	products, _ := project.GetAllProdutc()
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}
	product, _ := project.AddProduct()
	fmt.Println(product)
}
