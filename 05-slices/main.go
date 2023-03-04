package main

import "fmt"

func main() {
	fruits1 := []string{"mango", "banana", "orange"}
	fruits2 := []string{"apple", "lemon"}
	fruits2 = append(fruits2, "watermelon")
	fruits1 = append(fruits1, fruits2...)
	fmt.Println(fruits1)
	fruits1[3] = "khuldi"
	fruits1 = append(fruits1[:2], "lemon")
	fmt.Println(fruits1)
	fmt.Println(fruits1[3:])
	fruits3 := fruits1[2:4]
	fmt.Println(fruits3)

}
