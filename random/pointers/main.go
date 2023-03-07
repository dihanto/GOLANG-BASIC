package main

import "fmt"

func ChangeName(ptr *string) {
	*ptr = "Kurniadi"
}

func main() {
	var p *int
	var name string = "Kurniawan"
	x := 23
	p = &x
	fmt.Println(p)
	fmt.Println(*p)
	*p = 06
	fmt.Println(*p)
	fmt.Println(x)
	fmt.Printf("nama lama saya adalah %s, ", name)
	ChangeName(&name)
	fmt.Printf("nama baru saya adalah %s", name)
}
