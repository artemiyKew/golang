package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello")
	// arr := []int{2, 1, 2, 1, 0, 1, 2}
	fmt.Println(isPalindrome("race a car"))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	newString := ""
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		if (s[i] >= 'a' && s[i] >= 'z') || (s[i] >= '0' && s[i] >= '9') {
			newString += string(s[i])
		}
	}
	rString := ""
	fmt.Println(newString, rString)
	for i := len(newString) - 1; i >= 0; i-- {
		// fmt.Println(newString[i])
		rString += string(newString[i])
	}
	fmt.Println(newString, rString)

	return newString == rString
}
