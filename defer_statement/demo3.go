package deferstatement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("ürün kaydedildi :", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("loglandı")
}

func Demo3() {
	p := product{productName: "laptop", unitPrice: 5000}
	defer p.save() // önceki son halini alıyor
	p = product{productName: "mouse", unitPrice: 45}
	fmt.Println("başarılı")
	// defer i nereye koydugundan emin olman lazm
	//en sona koy
}
