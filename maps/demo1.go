package maps

import "fmt"

func Demo1() {
	sozluk := make(map[string]string)
	sozluk["table"] = "masa"
	sozluk["book"] = "kitap"
	sozluk["pencil"] = "kalem"

	fmt.Println(sozluk)
	fmt.Println(sozluk["book"])
	delete(sozluk, "book")
	fmt.Println("Eleman sayısı :", len(sozluk))

	// deger := sozluk["table"]
	deger, varMi := sozluk["table"] //2. parametre sözlükteki value değerini araştırıyor. Varsa true yoksa false döndürür.
	fmt.Println(deger)
	fmt.Println("Listede olma durumu :",varMi)

	// make kullanmadan map oluşturma
	sozluk2 := map[string]string{"glass":"bardak","microphone":"mikrofon"}
	fmt.Println(sozluk2)
	sozluk2["horse"] = "at"
	fmt.Printf("sozluk2: %v\n", sozluk2)
}
