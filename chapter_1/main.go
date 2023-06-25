/*
Первая программа
на языке Go
*/
package main // определение пакета для текущего файла
import "fmt" // подключение пакета fmt

// определение функции main
func main() {
	fmt.Println("Hello world!") // вывод строки на консоль

	var i int
	var a, b, c string
	var hello string = "Hello"
	var (
		name    string = "Vasya"
		surName string = "Petrov"
	)

	age := 10
	fmt.Println(i, a, b, c, hello, name, surName, "age =", age)

	TestDataTypes()
	ArrayTest()
	Conditions()
	Loops()
	Functions()
	Slices()
	Maps()
}
