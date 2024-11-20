package mapsingo

import (
	"fmt"
	"strings"
)

// Word frequency
func WordFrequency() {
	line := "hey this is a test for map. This test will check how many time each word is added in string. So I am going to repeat a few words please Ignore them they are just for the sake of this test."

	wordCount := make(map[string]int)
	sentence := strings.Fields(line)
	for _, words := range sentence {
		words = strings.ToLower(words)
		wordCount[words]++
	}
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// key check
func KeyCheck() {
	wordCheck := map[string]int{
		"empId":        1,
		"contact":      1234567,
		"country_code": 91,
	}
	value, ok := wordCheck["contact"]
	if ok {
		fmt.Printf("the value is %d\n", value)
	} else {
		fmt.Println("value Doesn't exist")
	}
}

// Reverse Map
func MapReverse() {
	amap := map[string]int{
		"empId":        1,
		"contact":      1234567,
		"country_code": 91,
	}
	reversedMap := make(map[int]string)
	for key, value := range amap {
		fmt.Printf("key is %s and value is %d\n", key, value)
		reversedMap[value] = key
	}
	fmt.Println("")
	fmt.Println("The Reversed values are :-")
	for rkey, rValue := range reversedMap {
		fmt.Printf("key is %d and its value is %s\n", rkey, rValue)
	}
}

// character frequency
func CharFrequency() {
	line := "hey this is a test for map. This test will check how many time each word is added in string. So I am going to repeat a few words please Ignore them they are just for the sake of this test."

	charCount := make(map[rune]int)
	// changedVaalue := strings.ToLower(line)
	for _, char := range line {
		charCount[char]++
	}
	fmt.Println("")
	fmt.Println("character frequency")

	for char, count := range charCount {
		fmt.Printf("The char is %q and its count is %d\n", char, count)
	}
}
