package error_handling
//Custom exception yazmak
import "fmt"

type borderException struct {
	parameter int
	message   string
}
//Bir struct'a error implemente edersen, Error görevi görebiliyor.

func (b *borderException) Error() string { //Error interfacesini implemente etmiş oluyorsun
	return fmt.Sprintf("%d---%s", b.parameter, b.message)
}

func TahminEt2(tahmin int) (string,error) {
	
	if tahmin <= 1 || tahmin >= 100 {
		return "", &borderException{tahmin, "Sınırların dışında kaldı."}  // & işareti pointer olarak kullanılıyor.
	}
	return "Bildiniz", nil
}

// & demek bir objenin veya val'ın ramdaki memory adresidir.
// * bir val'ın memory adresini tutar ve çözer (& operatörünün karşılığıdır)
//book_ptr := &new_book
//book_ptr.title = "Games of Thrones"