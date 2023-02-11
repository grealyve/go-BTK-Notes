package loops

import "fmt"

func Demo4() {
	var sayi int
	fmt.Println("Bir sayı giriniz : ")
	fmt.Scanln(&sayi)
	var durum bool = true
	
	if sayi > 0 {
		for i := 2; i < sayi; i++ {
			if sayi % i == 0 {
				durum = false
				break
			}
		}
	}else {
		fmt.Println("Lütfen pozitif bir sayı giriniz.")
	}
	
	if durum {
		fmt.Println("Sayı asaldır.")
	}else {
		fmt.Println("Sayı asal değil.")
	}
}