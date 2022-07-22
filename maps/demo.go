package maps

import "fmt"

// javadaki sözlük yapısı gibi
// eşlemek anlamında
func Demo() {
	// key -value anahtar deger mimarisinde kurulmuş
	sozluk := make(map[string]string)
	sozluk["table"] = "masa"
	sozluk["book"] = "kitap"
	sozluk["pencil"] = "kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("eleman sayısı : ", len(sozluk))
	fmt.Println("sözlük : ", sozluk)
	delete(sozluk, "book")
	fmt.Println("eleman sayısı : ", len(sozluk))

	deger, varMı := sozluk["pencil"]
	fmt.Println(deger)
	fmt.Println("listede olma durumu : ", varMı)

	sozluk2 := map[string]string{"glass": "bardak", "hey": "merhaba"}
	fmt.Println(sozluk2)
}
