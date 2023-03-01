package main

import "fmt"

func main() {
	currentYear := 2023
	//age := currentYear - 2009
	/*	if age < 17 {
			fmt.Println("Your age is not prohibitted to create Driver License")
		} else {
			fmt.Println("You are eligible to create Driver License")
		}*/
	// optimized logic
	message := "You are eligible to create Driver License"
	if age := currentYear - 2000; age < 17 {
		message = "Your age in not eligible to create Driver License"
	}
	fmt.Println(message)
}
