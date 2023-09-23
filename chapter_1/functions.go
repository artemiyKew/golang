package main

import "fmt"

func Functions() {
	fmt.Println("\nФункции и параметры!")
	// fmt.Println(add(5, 2))
	sum(1, 2, 3, 4, 5, 6, 7)
	r, err := returnValue()
	fmt.Println(r, err)

	f := add

	fmt.Println(f(3, 4, sum))
}

func sum(numbers ...int) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("sum =", sum)
}

func add(x, y int, opertaion func(...int)) int {
	return x + y
}

func returnValue() (int, string) {
	return 404, "Error"
}

/*

func название_функции (параметры_функции) возвращаемое_значение {

}

*/
