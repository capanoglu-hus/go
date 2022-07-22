package stringfunctions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "go programlama dili"
	fmt.Println(s.HasPrefix(isim, "go")) // ön ek
	fmt.Println(s.HasSuffix(isim, "ma")) // arka ek
	// boolean degerlerinde kullanılıyyor
	fmt.Println(s.Index(isim, "og"))
	harfler := []string{"p", "r", "o", "g", "r", "a", "m", "l", "a", "m", "a"}
	sonuc := (s.Join(harfler, " * "))
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "*", "+", -1))
	//28811111,020221,5000-tc-tarih-tutar
	fmt.Println(s.Split(sonuc, "-"))
	fmt.Println(s.Repeat(sonuc, 2)) //5tane kopyalasını yazıp verir

}
