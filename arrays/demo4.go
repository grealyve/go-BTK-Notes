package arrays

import "fmt"

func Demo4() {
	var numbers [2][3]int

	numbers[0][0] = 1
	numbers[0][1] = 2
	numbers[0][2] = 3
	numbers[1][0] = 4
	numbers[1][1] = 5
	numbers[1][2] = 6

	// fmt.Println(numbers[1][1])

	for x := 0; x < len(numbers); x++ {
		for i := 0; i < len(numbers[x]); i++ {
			fmt.Println(numbers[x][i])
		}
	}
}
