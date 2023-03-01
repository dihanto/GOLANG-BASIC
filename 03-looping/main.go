package main

import "fmt"

func main() {
	// first way of looping
	for i := 1; i < 6; i++ {
		fmt.Println("Number", i)
	}
	// second way of looping
	i := 1
	for i < 6 {
		fmt.Println("Number", i)
		i++
	}
	// third way of looping
	j := 1
	for {
		fmt.Println("Number", j)
		j++
		if j == 6 {
			break
		}
	}
	// 0 1 2 3 4
	// 1 2 3 4
	// 2 3 4
	// 3 4
	// 4
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, "")
		}
		fmt.Println()
	}
outerloop:
	for i := 0; i < 3; i++ {
		fmt.Println("Outerloop:", i)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerloop
			}
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
