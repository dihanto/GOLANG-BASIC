package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type Student struct {
	Name, Gender, Major string
	Age                 int
}

var std []Student

func main() {

	for {
		if Password() {
			Input()
		}
	}
}

func menu() int {
	var input int

	fmt.Println("1. Create Student")
	fmt.Println("2. Show Students")
	fmt.Println("3. Exit")
	fmt.Println("Enter the number to choose menu :")

	fmt.Scan(&input)

	return input
}

func Password() bool {
	var password string

	fmt.Println("Please enter the password to use this app")
	fmt.Scan(&password)
	return password == "jangkrik"
}
func Input() {
	for {
		input := menu()
		if input == 1 {
			var name, major, gender string

			name = strings.ToUpper(InputString("What is the name of student?"))
			age, _ := strconv.Atoi(InputInt("What is the age of student?"))
			aged := int(age)
			major = strings.ToUpper(InputString("What is the major of student?"))
			gender = GenderInput()

			student := Student{
				Name:   name,
				Age:    aged,
				Major:  major,
				Gender: gender,
			}
			std = append(std, student)

			db := initDatabaseConnection()
			query := "INSERT INTO student (name, age, major, gender) VALUES ($1,$2,$3,$4)"
			_, err := db.Exec(query, student.Name, student.Age, student.Major, student.Gender)
			if err != nil {
				log.Println(err)
			}

		} else if input == 2 {
			if len(std) == 0 {
				fmt.Println("No students found")
			} else {
				for i := 0; i < len(std); i++ {
					fmt.Printf("Name : %s\nAge : %d\nMajor : %s\nGender : %s\n", std[i].Name, std[i].Age, std[i].Major, std[i].Gender)
					fmt.Println("-------------------")
				}
			}
		} else if input == 3 {
			os.Exit(1)
		}
	}
}
func InputString(msg string) (result string) {

	for {
		var input string
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(msg)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
		}
		input = strings.TrimSpace(input)

		if len(input) == 0 {
			fmt.Println("input is empty, please type something")
			continue
		}

		_, err = strconv.Atoi(input)
		if err == nil {
			fmt.Println("input not valid, please read the instruction carefully.")
			continue
		} else if input == "true" {
			fmt.Println("input not valid, please read the instruction carefully.")
			continue
		} else if input == "false" {
			fmt.Println("input not valid, please read the instruction carefully")
			continue
		}

		result = input
		break
	}
	return result
}
func InputInt(msg string) string {
	var input string
	// var err error
	for {
		fmt.Println(msg)
		fmt.Scan(&input)

		_, err := strconv.ParseInt(input, 10, 32)

		if err != nil {
			fmt.Println("input not valid, please enter the age of student")
			continue
		}
		break
	}
	return input
}

func GenderInput() string {
	for {
		data := InputString("What is the gender of student :")
		if data == "male" {
			return fmt.Sprintln("Male")
		} else if data == "female" {
			return fmt.Sprintln("Female")
		} else {
			return "Please type male/female"
		}
	}
}

func initDatabaseConnection() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=cli sslmode=disable")
	if err != nil {
		fmt.Sprintln("Failed to connect database")
	}
	return db
}
