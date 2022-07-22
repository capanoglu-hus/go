package loops

import "fmt"

func Demo() {

	var metin string = "heyy"

	// for döngüsü farklı şekilde de yazılabiliyor
	i := 1
	// infinite loop  calışırsa sonsuz döngüye girer
	//her çalıştıgında i yi de arttırman lazım
	for i <= 5 {
		fmt.Println(metin)
		i = i + 1
	}
	fmt.Println("bitti")
	// while aynı yapıda döngü kurabilirsin

}
