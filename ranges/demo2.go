package ranges

import "fmt"

//1-10 arası tek sayıları toplayan fonksiyon
func Demo2()  {
	sayilar :=[]int{1,2,3,4,5,6,7,8,9,10}
	var toplam int
	for _, sayi := range sayilar{
		if sayi % 2 != 0{
			toplam += sayi
		}
	}
	fmt.Println("Toplam :",toplam)
}