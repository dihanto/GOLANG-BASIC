package main

import "fmt"

func main() {
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(evenNumbers(numbers...))
	var isPalindrome = func(word string) bool {
		var reversedWord string
		for i := len(word) - 1; i >= 0; i-- {
			reversedWord += string(word[i])
		}
		return word == reversedWord

	}("malam")
	fmt.Println(isPalindrome)

}
