package pointer

import "fmt"

//referans tip
func Demo(sayi *int) { //bellekteki adres yüzünden olacak
	*sayi = *sayi + 1
	fmt.Println("demodaki sayı = ", sayi)
}
