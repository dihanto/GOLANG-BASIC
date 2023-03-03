package main

import "fmt"

func main() {
	var alas float32 = 4
	var tinggi float32 = 6
	luas := 0.5 * alas * tinggi

	fmt.Println(luas)

	//menghitung luas banyak segitiga
	var jml float32
	fmt.Scanln(&jml)

	hasil := luas * jml

	fmt.Println(hasil)
}
