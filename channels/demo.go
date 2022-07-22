package channels

import (
	"fmt"
	"time"
)

func CiftSayilar(CiftSayiCn chan int) {
	toplam := 0
	for i := 0; i < 10; i += 2 {
		toplam = toplam + i
		fmt.Println("cift sayı toplama çalısıyor")
		time.Sleep(1 * time.Second)
	}
	CiftSayiCn <- toplam //en sonunda channel ile karşilaşıyor olunacak
	//işlemin bittigini göstermek için kullanıyor
}

func TekSayilar(TekSayiCn chan int) {
	toplam := 0
	for i := 1; i < 10; i += 2 {
		toplam = toplam + i
		fmt.Println("tek sayı toplama  çalısıyor")
		time.Sleep(1 * time.Second)
	}
	TekSayiCn <- toplam //
	// önemliiiii
	//channellarla asenkron çalışırdı.
	// hem de onların yaşam döngüsünü kontrol ediyoruz.
}
