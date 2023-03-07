package main

import (
	"fmt"
	"reflect"
)

func main() {
	type minute = uint
	type second = uint
	var hour minute = 60
	var minutes second = 60
	fmt.Println(reflect.TypeOf(hour))
	result := minutes * hour

	fmt.Println(result)
}
