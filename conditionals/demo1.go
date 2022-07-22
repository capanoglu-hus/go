package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("para yetersiz")
	}

	if cekilmekIstenen <= hesap {
		fmt.Println("para hazırlanıyor ")
		hesap = hesap - cekilmekIstenen
	}

	//sprintf kullanımı önemlii
	fmt.Println("bitti . Hesaptaki para :" + fmt.Sprintf("%v", hesap)) //Sprintf ile hesap degiskeni formatlanıyor
	// aynı ama farklı şekilde kullanım
	fmt.Printf("bitti . hesaptaki para : %v", hesap)
}
