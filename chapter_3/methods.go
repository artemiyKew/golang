package main

import "fmt"

/*

!	func (имя_параметра тип_получателя) имя_метода (параметры) (типы_возвращаемых_значений) {
!		тело_метода
!	}

*/

func (l library) print() {
	for _, val := range l {
		fmt.Println(val)
	}
}

func (p *person) eat() {
	fmt.Println(p.name, "eat")
}

func (p *person) print() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
}

func (p *person) updateAge(newAge int) {
	fmt.Println("before", p.age)
	p.age = newAge
	fmt.Println("after", p.age)
}

func Methods() {
	fmt.Println("Methods!")

	lib := library{"Book1", "book2", "book3"}
	lib.print()
	kylie := person{
		name: "kylie",
		age:  30,
		contactInfo: contact{
			email: "kylie@gmail.com",
			phone: "89999999999",
		},
	}
	kylie.print()
	kylie.eat()
	kylie.updateAge(33)
}
