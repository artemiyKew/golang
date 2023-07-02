package main

import "fmt"

func changeValue(x *int) {
	*x = (*x) * (*x)
}

func createPointer(x int) *int {
	p := new(int)
	*p = x
	return p
}

func Pointers() {
	fmt.Println("Pointers!")
	x := 10
	p := &x
	*p = 15
	fmt.Println(p, " ", *p)

	var pf *float64
	if pf != nil {
		fmt.Println("Value:", *pf)
	}
	var pNew = new(int)
	fmt.Println(pNew, *pNew)
	*pNew = 8
	fmt.Println(pNew, *pNew)

	x = 5
	fmt.Println("Old x: ", x)
	changeValue(&x)
	fmt.Println("New x: ", x)

	p1 := createPointer(23)
	p2 := createPointer(321)

	fmt.Println(p1, p2, *p1, *p2)
}
