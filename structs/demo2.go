package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() { //Fonksiyonu class ile ilişkilendirmek için bu şekilde yapıyorsun.
	fmt.Println("Eklendi :", c.firstName)
}

func (c customer) update() {
	fmt.Println("Güncellendi :", c.firstName)
}

func Demo2() {
	c := customer{firstName: "Engin", lastName: "Yanyatan", age: 39}
	c.save()
	c.update()
}
