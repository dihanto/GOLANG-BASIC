package main

import (
	"fmt"

	"github.com/dihanto/GOLANG-BASIC/01-GO-INTRODUCTION/data"
)

func main() {

	fmt.Println(data.FullName)
	fmt.Println(data.Age)
	fmt.Println(data.IsMarried)

	fmt.Printf("Nama saya %s umur saya %d tahun 8 bulan \n", data.FullName, data.Age)

	// var weight, tall int = 54, 168
	var (
		weight int
		tall   int
		city   string
	)
	weight, tall, city = 54, 168, "Tasikmalaya"
	fmt.Printf("Berat saya %d, tinggi saya %d dan saya berasal dari %s ", weight, tall, city)

	_, _, _ = weight, tall, city

}
