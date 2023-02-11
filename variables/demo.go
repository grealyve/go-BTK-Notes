package variables
//go mod init module_name 

import "fmt"

func Demo(){
	
	fmt.Println("Hello")
	fmt.Print("Word!")
	var metin string = "Merhaba Türkiye!"
	fmt.Println(metin)

	var kdv int = 14
	fmt.Println(kdv)

	var kdv2 float32 = 0.5
	fmt.Println(100+100*kdv2)

	kdv3 := 25
	fmt.Printf("Veri tipi : %T\n", kdv3) //%T type demek oluyor. Printf de formatlamaya yarıyor.

	var durum bool = true
	var metin1 string = "Engin"
	var metin2 string = "Ali"
	
	durum = metin1 == metin2
	fmt.Println(durum)
	fmt.Println(!durum)
}