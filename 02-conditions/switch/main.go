package main

import "fmt"

func main() {
	score := 4
	/*	switch score {
		case 8:
			fmt.Println("Perpect")
		case 7:
			fmt.Println("Awesome")
		default:
			fmt.Println("not Bad")
		}*/
	// relational operator
	switch {
	case score == 8:
		fmt.Println("Perpect")
	case (score > 3) && (score < 8):
		fmt.Println("Not Bad")
		fallthrough
	case score < 5:
		fmt.Println("It's okay, but please study harder")
	default:
		{
			fmt.Println("Study harder")
			fmt.Println("You need to learn more")
		}
	}
}
