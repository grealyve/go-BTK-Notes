package defers

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc() //Defer func. en aşşağıya koysan stack içine atmaz çalıştırmazdı.

	if sayi%2 == 0 {
		return "Çift sayıdır." //return sadece değeri döndürme olduğu için konsolda önce "Bitti" yazıyor.
	}

	if sayi%2 != 0 {
		return "Tek sayıdır." //return çalıştığı anda fonksiyonun en altına(süslü paranteze) imleç gider ve biter.
	}
	
	return ""

}

func Test() {
	sonuc := Demo2(9)
	fmt.Println(sonuc)
}

func DeferFunc()  {
	fmt.Println("Bitti")
}