package functions
//Variadic functions parametre sayısı belli değilken kullanılan bir fonksiyon türüdür.
//Sayıları array bazlı olarak alıyor. Array mantığında yaz fonksiyonunu.
func ToplaVariadic(sayilar ...int) int { 
	toplam := 0
	for i:=0;i < len(sayilar);i++{
		toplam += sayilar[i]
	}
	return toplam
}