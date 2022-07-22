package structs

import "fmt"

// class yapısı gibidir
// classla bellekte tutulma hali farklıdır
//deger alarak çalışır int gibi

func Demo() {
	fmt.Println(product{"bilgisayar", 500, "xyz", 12})
	fmt.Println(product{"klavye ", 300, "abc", 12})
	fmt.Println(product{name: "klavye ", unitPrice: 300})
	//özelliklerin hepsini yazmak istemezsen isimleriyle tek tek yaz
	// bir obje olmuş oldu
}

type product struct { // product adında bir structım var
	name         string  //ad
	unitPrice    float64 //fiyat
	brand        string  //marka
	discountRate int     // indirim oranı
	// buradki özelliklerin hepsini yazman lazım

}
