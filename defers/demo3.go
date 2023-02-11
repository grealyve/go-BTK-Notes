package defers

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) Save() {
	fmt.Println("Ürün kaydedildi :", p.productName)
	defer Log()
}

func Log()  {
	fmt.Println("Loglandı")
}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: 5000}
	defer p.Save()
	p = product{productName: "Mouse", unitPrice: 54}  //defer değerlerin en güncel haliyle değil, defer'den önceki p halini ekler.	

	fmt.Println("İşlem başarılı")
}
