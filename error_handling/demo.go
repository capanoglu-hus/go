package errorhandling

import (
	"fmt"
	"os"
)

// hata yakalama
// uygulamada oluşan hataları yakalayıp ona göre aksiyon olur
// 0'a bölünme istisnası -- payda 0 olursa tanımsız olur kullanıcan gelene bagllı
// veri kaynagına baglananama olabilr
// hatayı kontrol altına alma

func Demo() {
	f, err := os.Open("demo.txt") //os go tarafından oluyor
	// varsa dosyayı yoksa err dönüyor
	// dosya bulunursa err'in degeri nil degeri
	if err != nil { // yani hata varsa
		// type assertion oluşan hata türü
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("dosya bulunamadı :", pErr.Path)
			return
		} else {
			fmt.Println("dosya bulunamadı")
			return
		}

	} else {
		fmt.Println(f.Name())
	}
}
