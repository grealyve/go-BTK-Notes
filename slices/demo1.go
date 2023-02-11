package slices

import "fmt"

//Arraylist'in gelişmiş hali. Java'da olduğu gibi aynı. Daha çok bu kullanılır.
func Demo1()  {
	names := make([]string, 3) //3 indexli bir array oluşturuyor ama sınırı aşınca kızıyor. Fakat append ile ekleyebilirsin.

	fmt.Println(names)
	names[0] = "Ali"
	names[1] = "Veli"
	names[2] = "Haci"
	names = append(names, "Salih")
	fmt.Println(names)
	
}