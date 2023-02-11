package loops
//Go dilinde While döngüsü yok!
import "fmt"

func Demo1() {
	var metin string = "Merhaba Türkiye!"
	i := 1
	for i<=5{
		fmt.Println(metin)	//Bu şekilde bitmeyen döngü oluşturursun.
	}

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(metin)
	// }
}