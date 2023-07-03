package main

import "fmt"

/*

type имя_структуры struct {
	поля_структуры
}

*/

type contact struct {
	email string
	phone string
}

type person struct {
	name        string
	age         int
	contactInfo contact // !можно просто contact, одноименно со структурой
}

type node struct {
	value int
	next  *node
}

func Structs() {
	fmt.Println("Structs!")

	var tom person = person{
		name: "alex",
		age:  10,
		contactInfo: contact{
			email: "Hezam228@gmail.com",
			phone: "88005553535",
		},
	}
	fmt.Println(tom.name, tom.contactInfo.email)
	undefined := person{}
	fmt.Println(undefined)

	var tomPointer *person = &tom
	tomPointer.age = 28
	fmt.Println(tom.age)
	(*tomPointer).age = 30
	fmt.Println(tom.age)

	bob := new(person)
	_ = bob

	first := node{value: 5}
	second := node{value: 10}
	third := node{value: 15}

	first.next = &second
	second.next = &third

	currentNode := &first
	// for currentNode != nil {
	// 	fmt.Println(currentNode.value)
	// 	currentNode = currentNode.next
	// }
	printNodeValue(currentNode)
}

func printNodeValue(n *node) {
	if n != nil {
		fmt.Println(n.value)
		printNodeValue(n.next)
	}
}
