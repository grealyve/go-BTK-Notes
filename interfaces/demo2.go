package interfaces

import "fmt"

type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

// Interface'i implemente etmek için fonksiyonları structlar için yazman yeterlidir.
func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}

// farklı dosyada interface implementasyonu dene
//tüm fonksiyonlar zorunlu mu
func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator) float64 {
	payementTotal := 0.0

	for _, v := range credits {
		payementTotal += v.Calculate() //Polymorphizm
	}

	return payementTotal
}

func Demo2() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 100000, address: "Angara"}
	credit2 := Mortgage{rate: 12, creditPaymentTotal: 50000, address: "İstanbul"}
	credit3 := Car{rate: 15, creditPaymentTotal: 60000, carInfo: "Mini-Cooper"}

	credits := []CreditCalculator{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)

	fmt.Println("Toplam ödeme :", total)
}
