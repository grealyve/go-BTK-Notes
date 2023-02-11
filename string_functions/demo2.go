package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Engin"
	fmt.Println(s.HasPrefix(isim, "Eng")) //ön eki var mı diye kontrol etmeni sağlar.
	fmt.Println(s.HasSuffix(isim, "in"))
	fmt.Println(s.Index(isim, "gi"))
	harfler := []string{"e", "n", "g", "i", "n"}
	sonuc := (s.Join(harfler, "*")) //seperator istiyor 2.parametre olarak

	fmt.Println(s.Replace(sonuc, "*","+",-1)) //kaç tanesini değiştireceğini int param. yerine yaz. -1 hepsi demek

	fmt.Println(s.Split(sonuc, "*")) //array döndürür. ayıracı bulamazsa değişiklik yapmaz tek elemanlı array döndürür.
	fmt.Println(s.Repeat(sonuc, 5))
}
