package main

import "fmt"

func ArrayTest() {
	fmt.Println("\nРабота с массивами!")
	var numbers [6]int
	var numbers1 [5]int = [5]int{1, 2, 3, 4, 5}
	numbers2 := [5]int{2, 3, 4, 5, 6}
	numbers3 := [...]int{3, 4, 5, 6, 7, 8}
	for i := 0; i < 6; i++ {
		numbers[i] = i
	}

	nums := make([]int, 3)
	for i := 0; i < 3; i++ {
		nums[i] = i * i
	}

	fmt.Println(nums)

	matrix := [2][2]int{{1, 2}, {1, 2}}
	fmt.Println(numbers, numbers1, numbers2, numbers3)
	fmt.Println(matrix)
}
