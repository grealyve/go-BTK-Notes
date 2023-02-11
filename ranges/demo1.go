package ranges

import "fmt"

func Demo1() {
	sehirler := []string{"Ankara", "İzmir", "İstanbul"}
	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}

	for _, sehir := range sehirler { // _ yerine i yazabilirsin indexi öğrenmek için falan.
		fmt.Println(sehir)			// range dediğin zaman sehirler'i benim için bi gez diyorsun.
	}
}
