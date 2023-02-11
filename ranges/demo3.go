package ranges

import "fmt"

//mapi range ilde dönmek
func Demo3() {
	mymap := map[string]string{"at": "horse", "crocodile": "timsah", "snake": "yılan"}
	mymap["kuş"] = "bird"

	for k, v := range mymap {
		fmt.Println(k, v)
	}
}
