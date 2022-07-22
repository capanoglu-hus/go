package deferstatement

import "fmt"

// bir fonksiyonun en sonunda çalıştıgını garanti ettigimiz yapı
// fonksiyon bittiginde başka bir fonksiyon çalışmasını istedigimizde kullanılıyor
// çalışmasını garanti ediyor
// loglama alt yapıları için kullanılıyor
//

func A() {
	fmt.Println("a fonksiyonu çalıştı")
}
func C() {
	fmt.Println("c fonksiyonu çalıştı")
}
func D() {
	fmt.Println("d fonksiyonu çalıştı")
}

func B() {
	defer C()
	defer D()
	fmt.Println("b fonksiyonu çalıştı")
	defer A() // üstünde ne olursa olsun bitiminden sonra çalışır

}
