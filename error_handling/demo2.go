package errorhandling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	aklimdakisayi := 30

	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı gir ")
	}

	if tahmin > aklimdakisayi {
		return "daha küçük bir sayi gir", nil
	}
	if tahmin < aklimdakisayi {
		return "daha büyük bir sayi gir", nil
	}
	return "bildiginz", nil
}

func Demo2() {
	mesaj, hata := TahminEt(40)
	fmt.Println(mesaj)
	fmt.Println(hata)
}
