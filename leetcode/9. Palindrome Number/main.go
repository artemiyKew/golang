package main

import (
	"fmt"
)

func main() {
	fmt.Println(plusOne([]int{9}))
}

func plusOne(digits []int) []int {
	if (digits[len(digits)-1]+1)%10 != 0 {
		digits[len(digits)-1] += 1
		return digits
	} else {
		for i := len(digits) - 1; i >= 0; i-- {
			if i == 0 && (digits[i]+1)%10 == 0 {
				digits[i] = 0
				digits = append([]int{1}, digits...)
				break
			}
			if (digits[i]+1)%10 == 0 {
				digits[i] = 0
			} else {
				digits[i] += 1
				break
			}
		}
	}
	return digits
}
