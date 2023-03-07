package main

import (
	"fmt"
)

func main() {
	city := "Jakarta"

	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, city[i])
	}
	bytesCity := []byte{74, 97, 107, 97, 114, 116, 97}
	for index, value := range bytesCity {
		fmt.Printf("index: %d, string: %s\n", index, string(value))
	}
	var name string = "Kurniawan"
	var age int = 22
	var bio string = fmt.Sprintf("My name is %s, I am %d years old, I am from %s", name, age, city)
	fmt.Println(bio)
}
