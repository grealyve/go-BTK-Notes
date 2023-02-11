package string_functions

//alias  (import as ... gibi)
import (
	"fmt"
	s "strings"
)

func Demo1() {
	isim := "Engin"
	fmt.Println(s.Count(isim, "e"))		//Fonk. açıklaması için ctrl + Space
	fmt.Println(s.Contains(isim, "a"))
	fmt.Println(s.Index(isim, "n")) //ilk gördüğünün indexini döndürür bulamazsa -1 döndürür
	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))


}
