package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func Concat(textString string, sep string) map[int]string {
	slices := []string{}
	keys := make(map[string]bool)

	result := map[int]string{}
	slices = strings.Split(textString, " ")

	fmt.Println(slices)
	for _, val := range slices {
		if _, value := keys[val]; !value {
			keys[val] = true
			result[strings.Count(textString, val)] = val
		}

	}
	eys := make([]int, 0, len(result))
	for k := range result {
		eys = append(eys, k)
	}
	sort.Ints(eys)

	return result
}

func main() {

	testInput := "Hello Hello Hello bye bye yes yes yes yes"
	testOutput := map[int]string{2: "bye", 3: "Hello", 4: "yes"}
	r2 := Concat(testInput, " ")
	fmt.Printf("Test: %v\n", testInput)
	fmt.Printf("Expected: %v\n", testOutput)
	fmt.Printf("Result: %v\n", r2)
	if reflect.DeepEqual(testOutput, r2) {
		fmt.Printf("OK\n\n")
	} else {
		fmt.Printf("FAIL\n\n")
	}

}
