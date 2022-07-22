package stringfunctions

//alias denir
import (
	"fmt"
	s "strings" //strings i s yaptıgım zaman da çalışacaktır
)

func Demo() {
	isim := "hasdasç"
	fmt.Println(s.Count(isim, "a")) // isimin içinde kaç tane g var diye say demek
	// case sensitive büyük küçük harf duyarlı
	fmt.Println(s.Contains(isim, "a")) // a harfi kelime içinde varsa true yoksa false dönüyor
	sonuc := s.Index(isim, "n")        // kaçıncı  indekste oldugunu döndürür ilk için geeçerli

	if sonuc != -1 {
		fmt.Println("a harfi var")
	} else {
		fmt.Println("a harfi yok")
	}
	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))
}
