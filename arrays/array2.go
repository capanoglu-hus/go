package arrays

import "fmt"

func Demo2() {
	var sehirler [5]string
	sehirler[0] = "ankara"
	sehirler[1] = "istanbul"
	sehirler[2] = "izmir"
	sehirler[3] = "adana"
	sehirler[4] = "mersin"

	fmt.Println(sehirler)
	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}
}
