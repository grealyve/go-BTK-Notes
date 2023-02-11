package slices

import "fmt"

func Demo2() {
	cities := []string{"Muğla", "İzmir", "Bursa"} //arka tarafta make yapıyor bunu yazınca.

	//Bir Slice'ı kopyalamak
	copyOfCities := make([]string, len(cities))
	fmt.Println(copyOfCities)
	copy(copyOfCities, cities)
	fmt.Println(copyOfCities)

	cities = append(cities, "İstanbul") //iki ayrı slice oluşuyor.
	fmt.Println(cities)
	fmt.Println(copyOfCities)

	fmt.Println(cities[1:3]) //parçalamak Python ile aynı
	fmt.Println(cities[:3])
	fmt.Println(cities[1:])

}
