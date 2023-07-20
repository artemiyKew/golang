package main

import "fmt"

type Number interface {
	~int64 | float64 // ~ - означает приближение типа
}

type CustomInt int64

func (ci CustomInt) isPositive() bool {
	return ci > 0
}

type Numbers[T Number] []T

func Generic() {
	// showSum()
	// showContains()
	// showAny()
	// unionInterfaceAndType()
	typeApproximation()
}

func showSum() {
	floats := []float64{1.0, 2.0, 3.0}
	ints := []int64{1, 2, 3}

	fmt.Println(sum(floats))
	fmt.Println(sum(ints))
}

func sum[V int64 | float64](nums []V) V {
	var sum V
	for _, num := range nums {
		sum += num
	}
	return sum
}

func showContains() {
	type Person struct {
		name     string
		age      int64
		jobTitle string
	}

	ints := []int64{1, 2, 3, 4, 5}
	fmt.Println("int:", contains(ints, 4))

	strings := []string{"Vasya", "Dima", "Katya"}
	fmt.Println("strings: ", contains(strings, "Katya"))
	fmt.Println("strings: ", contains(strings, "Sasha"))

	people := []Person{
		{
			name:     "Vasya",
			age:      15,
			jobTitle: "MediaSoft",
		},
		{
			name:     "Dima",
			age:      28,
			jobTitle: "VK",
		},
		{
			name:     "Katya",
			age:      25,
			jobTitle: "Yandex",
		},
	}
	fmt.Println("people:", contains(people, Person{
		name:     "Vasya",
		age:      15,
		jobTitle: "MediaSoft",
	}))
}

func contains[T comparable](elems []T, searchEl T) bool {
	for _, elem := range elems {
		if searchEl == elem {
			return true
		}
	}
	return false
}

func showAny() {
	show(1, 2, 3)
	show("test1", "test2", "test3")
	show([]int64{1, 2, 3}, []int64{4, 5, 6})
	show(map[string]int64{
		"first":  1,
		"second": 2,
	})
	show(interface{}(1), interface{}("string"), any(struct{ name string }{name: "vasya"}))
}

func show[T any](elems ...T) {
	for _, elem := range elems {
		fmt.Println(elem)
	}
}

func unionInterfaceAndType() {
	var ints Numbers[int64]
	ints = append(ints, []int64{1, 2, 3, 4, 5}...)

	floats := Numbers[float64]{1.0, 2, 5, 3, 5}

	fmt.Println(sumUnionInterface(ints))
	fmt.Println(sumUnionInterface(floats))
}

func sumUnionInterface[T Number](elems []T) T {
	var sum T
	for _, elem := range elems {
		sum += elem
	}
	return sum
}

func typeApproximation() {
	customInts := []CustomInt{0, 2, 3, 4, 5}
	castedInts := make([]int64, len(customInts))

	for idx, val := range customInts {
		if val.isPositive() {
			fmt.Println(val, true)
		}
		castedInts[idx] = int64(val)
	}
	fmt.Println(sumUnionInterface(castedInts))
	fmt.Println(sumUnionInterface(customInts))
}
