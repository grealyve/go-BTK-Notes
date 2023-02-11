package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	aklimdakiSayi := 51

	if tahmin <= 1 || tahmin >= 100 {
		return "", errors.New("1-100 arasında sayı giriniz")
	} else if tahmin > aklimdakiSayi {
		return "Daha küçük sayı giriniz.", nil
	}else if tahmin < aklimdakiSayi {
		return "Daha büyük sayı giriniz.", nil
	}

	return "Bildiniz", nil
}

func Demo2() {
	mesaj, hata := TahminEt(50)
	fmt.Println(mesaj)
	fmt.Println(hata)
	fmt.Println(TahminEt(-10))

}
