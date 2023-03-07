package variadic

import (
	"fmt"
)

func Print(names ...string) []map[string]string {
	var result []map[string]string
	for i, v := range names {
		key := fmt.Sprintf("stundent%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}

func Sum(numbers ...int) (result int) {
	for _, value := range numbers {
		result += value
	}
	return
}
