package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1900
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("Hesaptaki para yetersiz.")
	}

	if cekilmekIstenen <= hesap {
		fmt.Println("Para hazırlanıyor")
		hesap = hesap - cekilmekIstenen
	}

	// fmt.Println("Bitti. Hesaptaki para : " + fmt.Sprintf("%v",hesap))
	// fmt.Println("Bitti. Hesaptaki para : ", hesap)
	fmt.Printf("Bitti. Hesaptaki para : %v", hesap)
}