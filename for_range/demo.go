package for_range

import "fmt"

func Demo() {
	sehirler := []string{"ankara", "mersin", "adana"}
	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])

	}

	//range demek sehirleri benim için gez demek
	for i, sehir := range sehirler {
		fmt.Println(i)
		fmt.Println(sehir)

	}

}
