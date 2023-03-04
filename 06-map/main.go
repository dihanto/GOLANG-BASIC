package main

import (
	"fmt"
)

func main() {
	human := make(map[string][]string)

	human["name"] = append(human["name"], "kurniawan")
	human["name"] = append(human["name"], "ujang")
	human["gender"] = append(human["gender"], "man")
	// human := map [string]string {
	// "name" : "kurniawan",
	// "gender" : "man",
	// }

	// fmt.Println("name:", human["name"], "gender:", human["gender"])
	// n := make(map[string]string)
	// fmt.Println(n)
	// for key, value := range human {
	// fmt.Println(key, value)
	// }
	// delete(human, "name")
	// human["name"] = "kurniadi"
	fmt.Println(human)
	person := []map[string]string{
		{
			"name":   "kurniadi",
			"gender": "man",
		},
		{
			"name":   "ujang",
			"gender": "man",
		},
	}
	fmt.Println(person)
}
