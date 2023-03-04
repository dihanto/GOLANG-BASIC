package main

import "fmt"

func main() {
	var pedestal float32 = 4
	var tall float32 = 6
	wide := 0.5 * pedestal * tall

	fmt.Println(wide)

	//menghitung luas banyak segitiga
	var amount float32
	fmt.Scanln(&amount)

	result := wide * amount

	fmt.Println(result)
}
