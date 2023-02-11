package interfaces

import "fmt"

func dogrula(i interface{}) { //interface her şeyi kapsar o yüzden her şeyi atayabilirsin
	// sayi := i.(int) //sayının int olup olmadığını kontrol eder.
	sayi, ok := i.(int) //ok -> değeri atayıp atayamadığını kontrol etmek için onu yakalamak için. Yakalayamazsa sayi 0 olur.
	fmt.Println(sayi, ok)
}

func Demo3() {
	var sayi interface{} = "Engin" //tip dönüşümünü yapamıyor 0 false çeviriyor.
	dogrula(sayi)

	var sayi2 interface{} = 5
	dogrula(sayi2)
}
