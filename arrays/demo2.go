package arrays

import "fmt"

func Demo2() {
	var cities [5]string
	cities[0] = "Angara"
	cities[1] = "İstanbul"
	cities[2] = "İzmir"
	// cities[3] = "Muğla" //Böyle yazdırınca listede iki boşluk şeklinde gösterir ama aslında "" halinde listede tutulur
	cities[4] = "Bursa"
	fmt.Println(cities)

	for i := 0; i < 5; i++ {
		fmt.Println(cities[i])
	}
}