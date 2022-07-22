package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() { //müsteriyle save i birleştirmek istiyorum
	fmt.Println("eklendi : ", c.firstName) // metot olarak geldi
}

func (c customer) update() { //müsteriyle save i birleştirmek istiyorum
	fmt.Println("güncelleme : ", c.firstName) // metot olarak geldi
}

func Demo2() {
	c := customer{firstName: "husniye", lastName: "capan", age: 23}
	c.save()
	c.update()
}
