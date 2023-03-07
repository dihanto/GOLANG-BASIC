package main

import (
	"fmt"
	"strings"

	"github.com/dihanto/golang-basic/09-function/variadic"
)

func main() {
	printGreet("Airell", 23)
	var names = []string{"Airell", "Jordan"}
	var printMsg, length = greet("Heii", names)

	fmt.Println(printMsg, length)

	studentLists := variadic.Print("Airell", "Nanda", "Mailo", "Schannel", "Marco")

	fmt.Printf("%v\n", studentLists)
	var numbers = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(variadic.Sum(numbers...))
	result := variadic.Sum(3, 5, 5, 10, 100, 1000)
	fmt.Println(result)
}
func printGreet(name string, age int8) {
	fmt.Printf("Hello there! My name is %s ad I'm %d years old\n", name, age)
}
func greet(msg string, names []string) (result string, length int) {
	var joinStr = strings.Join(names, " ")
	result = fmt.Sprintf("%s %s", msg, joinStr)
	length = len(result)
	return
}
