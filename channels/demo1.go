package channels

//Gerçek hayatta sleep kullanamayacağın için channel mantığını kullanacaksın.

//Parametre olarak channel oluşturacaksın
func CiftSayilar(ciftSayiCn chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam += i
	}

	//Channel'i kapatma işlemi bu. Channel'in taşıyacağı değeri belirtmek aslında yaptığımız.
	ciftSayiCn <- toplam
}

//Channel bir kanal görevi görüyor ve bu işlemin bittiğini garanti etmeye yarıyor 
//sleep kullanmadan böyle yapılır.
func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam += i
	}

	tekSayiCn <- toplam
}
