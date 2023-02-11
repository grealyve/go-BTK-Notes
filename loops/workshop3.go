package loops

import "fmt"

func Demo5() {
	sayi1 := 220
	sayi2 := 284

	sayi1toplam := 0
	sayi2toplam := 0

	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			sayi1toplam += i
		}
	}

	for i := 1; i < sayi2; i++ {
		if sayi2%i == 0 {
			sayi2toplam += i
		}
	}

	if sayi1toplam == sayi2 && sayi2toplam == sayi1 {
		fmt.Printf("%v ve %v arkadaş sayılardır.", sayi1, sayi2)
	}else{
		fmt.Printf("%v ve %v arkadaş sayı değillerdir.", sayi1, sayi2)
	}
}