package main

import "fmt"

func main() {
	numbers := [4]int{1, 2, 3, 4}
	fmt.Println(numbers)
	names := [3]string{"Kurniawan", "Kurniadi", "asep"}
	fmt.Println(names)
	names[1] = "ujang"
	fmt.Println(names[1])
	for i := 0; i < 3; i++ {
		fmt.Println(i+1, names[i])
	}
	//array  multidimensional
	balances := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(balances[1][1])
	peoples := [4][2]string{
		{"nurdin", "pemalang"},
		{"dedi", "pekanbaru"},
		{"ahmad", "bogor"},
		{"ujang", "pekanbaru"},
	}
	// name : ujang
	// city : pekanbaru
	// ------------------
	for i := 0; i < len(peoples); i++ {
		fmt.Println("name :", peoples[i][0])
		fmt.Println("city :", peoples[i][1])
		fmt.Println("-----------------------")
	}
}
