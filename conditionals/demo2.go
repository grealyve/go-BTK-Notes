package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1900
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("Hesaptaki para yetersiz.")
	} else if cekilmekIstenen == hesap{
		fmt.Println("Para geliyoor!")
		fmt.Println("Dikkat paran bitti!")
	} else {
		fmt.Println("Para geliyoor!")
	}

}