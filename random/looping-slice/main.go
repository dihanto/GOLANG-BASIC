package main

import "fmt"

func main() {
	numbers := []int{12, 7, 22, 15, 9, 4, 17}

	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			fmt.Println(numbers[i], "adalah bilangan genap")
		}
	}
}
