package main

import "fmt"

func main() {
	for i := 0; i < 13; i++ {
		if i == 8 {
			break
		}
		fmt.Println("ini perulangan ke", i)
	}
	for i := 0; i < 13; i++ {
		if i == 8 {
			continue
		}
		fmt.Println("ini perulangan ke", i)
	}
}
