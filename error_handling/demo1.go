package error_handling

import (
	"fmt"
	"os"
)

//type assertion
func Demo1() {
	f, err := os.Open("demo1.txt") //varsa dosyayı, yoksa error dönüyor bu fonksiyon
	//hata yoksa err = nil olur, varsa nil den farklı bir şeydir.

	//Daha iyi anlamak için interfaces -> demo3
	if err != nil {
		//err.(...) bu error'un tipini belirlemek için.
		// * koymasan da olur ama pointer değerini alması için ekstra nesne oluşmaması için böyle yap.
		if pErr, ok := err.(*os.PathError); ok {	// eğer hatanın tipi path error ise ok = true olur ve pErr'e path error atanır.
			fmt.Println("Dosya bulunamadı : ", pErr.Path) // ; 'e kadar olan kısım atama yapma işlemi, ok true ise aşağısı çalışyıor.
			return
		} else {
			fmt.Println("Genel hata oluştu.")
			return
		}
	} else {
		fmt.Println(f.Name()) //Çalıştırılan kod main.go da çalıştığı için onun dizininde olması gerek.
	}
}

// 	BU KISIM TYPE ASSERTION OLARAK GEÇER.
// if pErr, ok := err.(*os.PathError); ok {
// 	fmt.Println("Dosya bulunamadı : ", pErr.Path)
// 	return
// } else {
// 	fmt.Println("Dosya bulunamadı : ")
// 	return
// }