package main

import "fmt"

func main() {
	var input int

	fmt.Scanln(&input)

	if input%2 == 1 {
		fmt.Println(input, "adalah bilangan ganjil")
	} else {
		fmt.Println(input, "adalah bilangan genap")
	}
}
