package conditionals

import "fmt"

func Demo3() {
	// var x, y, z int

	// if x > y && x > z {
	// 	fmt.Println("En büyük sayı : x")
	// }else if y > x && y > z{
	// 	fmt.Println("En büyük sayı : y")
	// }else {
	// 	fmt.Println("En büyük sayı : z")
	// }

	var sayi1, sayi2, sayi3 int = 10,5,18

	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2		
	}else if sayi3 > enBuyuk{
		enBuyuk = sayi3
	}

	fmt.Printf("En büyük sayı : %v", enBuyuk)
}