package arrays

import "fmt"

func Demo3() {
	numbers := [5]int{20, 30, 50, 40, 2}
	fmt.Println(numbers)
	// myarray := [...]int{1, 25, 6, 12, 6}
	greatestNumber := 0
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
		if numbers[i] > greatestNumber {
			greatestNumber = numbers[i]
		}
	}
	fmt.Printf("Listedeki en büyük sayı : %v dır.", greatestNumber)
}
