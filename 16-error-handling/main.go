package main

import (
	"errors"
	"fmt"
)

func main() {

	defer catchErr()

	// var number int
	// var err error
	// number, err = strconv.Atoi("123GH")

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(number)

	// }

	// number, err = strconv.Atoi("123")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(number)

	// }
	var password string
	fmt.Print("Enter password : ")
	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}

}

func validPassword(password string) (string, error) {

	if len(password) < 5 {
		return "", errors.New("password has to have more 4 characters")
	}
	return "Valid password", nil
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Applicaion running perpectly")
	}
}
