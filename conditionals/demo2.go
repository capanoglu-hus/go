package conditionals

import "fmt"

func demo2() {
	var hesap float64 = 1200
	var cekilmekIstenen float64 = 1100

	if cekilmekIstenen > hesap {
		fmt.Println("para yetersiz")
	} else if cekilmekIstenen == hesap {
		fmt.Println("para kalmadı hesapta")
	} else {
		fmt.Println("para hazırlanıyor ")
	}

	//if sart{ if sart} iç içe if yazabilirsin

	//sprintf kullanımı önemlii
	fmt.Println("bitti . Hesaptaki para :" + fmt.Sprintf("%v", hesap)) //Sprintf ile hesap degiskeni formatlanıyor
	// aynı ama farklı şekilde kullanım
	fmt.Printf("bitti . hesaptaki para : %v", hesap)
}
