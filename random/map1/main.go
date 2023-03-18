package main

import "fmt"

func main() {

	angka := map[int]string{
		1: "satu",
		2: "dua",
		3: "tiga",
		4: "empat",
	}
	delete(angka, 3)
	fmt.Println(angka)

	var input string
	fmt.Println("Masukan sebuah string : ")
	fmt.Scanln(&input)

	freq := make(map[rune]int)
	for _, char := range input {
		freq[char]++
	}
	fmt.Println(freq)

}
