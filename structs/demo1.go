package structs

import "fmt"
//Bildiğin class yapısı
func Demo1() {
	fmt.Println(product{"Bilgisayar", 500, "ACER"})
	//Eksik parametre verdiğinde kızar ama parametreleri isimli olarak verirsen kızmaz.
	fmt.Println(product{name: "Bilgisayar", unitPrice: 500})
	fmt.Println(product{name: "Bilgisayar"}) //Sıralı olarak vermen gerekmez.

}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
