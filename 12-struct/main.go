package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

type Employee struct {
	// Name     string
	// Age      int
	Person
	Division string
	Salary   float32
}

func (person Person) Greet() string {
	return fmt.Sprintf("Hello I am %s", person.Name)
}
func (person *Person) ChangeName(newName string) {
	person.Name = newName
}

func main() {
	var employee1 Employee
	employee1.Name = "John"
	employee1.Age = 12
	employee1.Division = "IT"
	employee1.Salary = 20000.00
	employee1.Person.ChangeName("Usman")
	fmt.Println(employee1.Person.Greet())
	fmt.Println(employee1)
	// employee2 := Employee{
	// 	Name:     "Doe",
	// 	Age:      22,
	// 	Division: "EHS",
	// 	Salary:   30000.00,
	// }
	fmt.Println(employee1.Name)
	// fmt.Println(employee2)
	employee3 := Employee{}
	fillStruct(&employee3)
	fmt.Println(employee3)
	employee4 := struct {
		person   Person
		Division string
		Salary   float32
	}{
		person: Person{
			Name: "Ujang",
			Age:  22,
		},
		Division: "Production",
		Salary:   22000,
	}
	fmt.Println(employee4)
	employees := []Employee{
		{
			Person: Person{
				Name: "udi",
				Age:  11,
			},
			Division: "IT",
			Salary:   2110.00,
		},
	}
	employees = append(employees, employee3)
	employees = append(employees, Employee{
		Person: Person{
			Name:    "Umar",
			Age:     32,
			Hobbies: []string{"Hiking", "Gaming"},
		},
		Division: "GA",
		Salary:   22103.00,
	})
	fmt.Println(employees)
	fmt.Println(employees[2].Person.Name)
	fmt.Println(employees[2].Person.Hobbies[1])
}
func fillStruct(employee *Employee) {
	employee.Name = "Kurniawan"
	employee.Age = 23
	employee.Division = "HR"
	employee.Salary = 23000.00
}
