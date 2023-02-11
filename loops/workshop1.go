package loops

import "fmt"

//Sayı tahmin oyunu
func Demo3()  {
	aklimdakiSayi := 80
	var tahmniEdilenSayi, counter int

	for tahmniEdilenSayi != aklimdakiSayi{
		fmt.Println("Bir sayı tahmin ediniz : ")
		fmt.Scanln(&tahmniEdilenSayi)

		if tahmniEdilenSayi == aklimdakiSayi{
			counter += 1
		}else if tahmniEdilenSayi > aklimdakiSayi {
			fmt.Println("Yanlış tahmin, aşşağı in")
			counter++
		}else if tahmniEdilenSayi < aklimdakiSayi{
			fmt.Println("Yanlış tahmin, yukarı çık")
			counter += 1
		}else {
			fmt.Println("Eksik ya da hatalı tuşladınız.")
		}
	}
	var basarim string
	if  counter > 0 && counter <= 3 {
		basarim = "Süper"
	}else if counter >= 4 && counter <= 6 {
		basarim = "Güzel"
	}else {
		basarim = "Fena Değil"
	}

	fmt.Printf("Bravo bildiniz, %v tahmin %v !!!!", counter, basarim)
}