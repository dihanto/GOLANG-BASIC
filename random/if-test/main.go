package main

import "fmt"

func main() {

	var belanja int
	fmt.Println("Sudah berapa kali anda belanja bulan ini?? ")
	fmt.Scanln(&belanja)

	if belanja >= 20 {
		fmt.Println("Anda berhak mendapatkan diskon")
	} else if belanja < 20 {
		fmt.Println("Maaf anda tidak berhak mendapatkan diskon")
	}

}
