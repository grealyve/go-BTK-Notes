package pointers

import "fmt"

//Bir variable'ın bellekteki değerini de değiştirmek için pointer kullan.
func Demo1(sayi *int) {
	*sayi = *sayi + 1
	fmt.Println("Demodaki sayı :", *sayi) //Bu satırda * koymazsan ramdaki değerini gösterir.

	//sayi2 := sayi1 yapsan bile sadece değerini kopyalacaktır. sayi1 değiştiğinde sayi2 değişmeyecektir. 
}
