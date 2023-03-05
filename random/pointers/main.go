package main

import "fmt"

func main() {
	var p *int
	x := 23
	p = &x
	fmt.Println(p)
	fmt.Println(*p)
	*p = 06
	fmt.Println(*p)
	fmt.Println(x)

}
