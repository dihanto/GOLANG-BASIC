package main

import "fmt"

func main() {
	// var firstNumber int = 4
	// var secondNumber *int = &firstNumber
	// fmt.Println("firstNumber (value) :", firstNumber)
	// fmt.Println("firstNumber (memori address) :", &firstNumber)

	// fmt.Println("secondNumber (value) :", *secondNumber)
	// fmt.Println("secondNumber (memroi address) :", secondNumber)
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("secondNumber (value) :", *secondPerson)
	fmt.Println("secondNumber (memori address) :", secondPerson)

	fmt.Println("#################################################################")

	*secondPerson = "Doe"

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("secondNumber (value) :", *secondPerson)
	fmt.Println("secondNumber (memori address) :", secondPerson)

	changeName(&firstPerson)
	fmt.Println(firstPerson)

}
func changeName(name *string) {
	*name = "udin"
}
