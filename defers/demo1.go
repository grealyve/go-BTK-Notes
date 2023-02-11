package defers

import "fmt"

//bir fonksiyon bittikten sonra başka bir fonksiyonun çalışmasını garanti eder hata alırsan falan diye.
// Loglama, bağlantı kapama...

func A()  {
	fmt.Println("A fonksiyonu çalıştı.")
}

func C()  {
	fmt.Println("C fonksiyonu çalıştı.")
}

func B()  {		//defer fonk. çalışması için önce içindeki fonksiyon işini bitirmesi lazım.
	defer A()	//Defer'leri stack içine atıyor first in last out şeklinde çalışıyor.
	defer C()
	fmt.Println("B fonksiyonu çalıştı.")
}