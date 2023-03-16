package main

import (
	"fmt"

	"github.com/dihanto/golang-basic/practice/bank"
	"github.com/dihanto/golang-basic/practice/triangle"
)

func main() {
	triangle1 := triangle.Triangle{
		Base: 5.0,
		Tall: 8.0,
	}
	fmt.Println(triangle1.Wide())

	account1 := bank.BankAccount{
		Name:          "Yayan",
		AccountNumber: 2031239310,
		Balance:       5000000,
	}
	fmt.Println(account1.Balance)
	account1.Withdraw(100000)
	fmt.Println(account1.Balance)
	fmt.Println(account1.Balance)
}
