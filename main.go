package main

//go mod init module_name
import (
	"fmt"
	"golesson/project"
)

//Burda modülün içindeki bir paketi çağırdık.
//Modül, bir çok paketi barındıran üst paket gibi düşünülebilir.

func main() {
	// variables.Demo()
	// fmt.Print()
	// conditionals.Demo2()
	// conditionals.Demo3()
	// loops.Demo5()
	// arrays.Demo4()
	// slices.Demo2()
	// functions.SelamVer("Ali")
	// var sonuc = functions.Topla(2,6)
	// sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(10,6)

	// fmt.Println("Toplam :",sonuc1)
	// fmt.Println("Çıkarım :",sonuc2)
	// fmt.Println("Çarpım :",sonuc3)
	// fmt.Println("Bölüm :",sonuc4) //istemediğin bir return değeri yerine _ koyunca sıkıntı çıkmaz.

	// fmt.Println(functions.ToplaVariadic(1, 5, 2, 8, 4))
	// sayilar := []int{1, 5, 2, 8, 4}
	// fmt.Println(functions.ToplaVariadic(sayilar...)) //array gönderdiğin zaman onun variadic olduğunu belirtmek için ... konur sonuna.

	// maps.Demo1()
	// ranges.Demo3()

	// sayi := 20                          //Fonksiyon, bu değeri kopyalayıp kullanır. Bellektekini kullanmak istiyorsan * koy.
	// pointers.Demo1(&sayi)               //Sayıyı normalde yolladığında sadece demodaki değeri değişir bellekteki "sayı" değeri değişmez.
	// fmt.Println("Maindeki sayı:", sayi) //Bellekteki değeri de değiştirmek için pointerları kullanman gerek.
	// sayilar:=[]int{1,2,3}		//ÇOKOMELLİ!!! Array bazlı yapılar adresi ile çalışır. Yani bir şey değiştirdiğinde gerçeğinde de değişir.
	// pointers.Demo2(sayilar)
	// fmt.Println("Maindeki sayı:", sayilar[0])

	// structs.Demo2()

	//Normalde yazılım dilleri yukarıdan aşşağıya doğru çalıştırır fonksiyonları teker teker.
	//Aşağıdaki 2 fonksiyonu eş zamanlı çalıştırmak için başına go yazıyorsun ve main fonksiyonu bunları
	//farklı bir thread'e aktarmış oluyor ama kendi işi bittiği için kapatıyor ve hiçbir şey yazmıyor.
	// go goroutines.CiftSayilar()
	// go goroutines.TekSayilar()
	// time.Sleep(5 * time.Second)
	// fmt.Println("Main bitti")

	//Bu channellar sayı gibi ama değillerdir. Sadece içinde değer taşırlar.
	// ciftSayiCn := make(chan int)
	// tekSayiCn := make(chan int)
	// go channels.CiftSayilar(ciftSayiCn)
	// go channels.TekSayilar(tekSayiCn)

	//Channel değerlerini bir int gibi variable'a atadık burada.
	// ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn

	// carpim := ciftSayiToplam * tekSayiToplam
	// fmt.Println("Çarpım :",carpim)

	// interfaces.Demo2()

	// defers.Demo3()

	// error_handling.Demo1()

	// interfaces.Demo3()

	// error_handling.Demo2()

	// fmt.Println(error_handling.TahminEt2(102))

	// string_functions.Demo2()

	// restful.Demo2()

	// project.AddProduct()
	// project.GetAllProducts()

	products, _ := project.GetAllProducts()
	for _, v := range products {
		fmt.Println(v)
	}
}
