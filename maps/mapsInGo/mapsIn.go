package mapsingo

import (
	"encoding/json"
	"fmt"
	"log"
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

// Sum of values
func SumValues() {
	bill := map[string]int{
		"Grocery": 100,
		"Petrol":  2199,
		"Movie":   956,
	}
	x := 0
	fmt.Println("")
	fmt.Println("Sum of values")
	for key, value := range bill {
		fmt.Printf("key is %s and value is %d\n", key, value)
		x += value
	}
	fmt.Printf("the sum of all the values is %d\n", x)
}

// Merge Two Maps
func MergeMaps() {
	amap := map[string]int{
		"empId":        1,
		"contact":      1234567,
		"country_code": 91,
	}
	empBill := map[string]int{
		"Grocery": 100,
		"Petrol":  2199,
		"Movie":   956,
	}
	for key, value := range empBill {
		amap[key] = value
	}
	fmt.Println("")
	fmt.Println("merged map is below")
	for key, value := range amap {
		fmt.Printf("key is %s and its value is %d\n", key, value)
	}
}

// Find Keys by value
func KeyByValue() {
	empBill := map[string]int{
		"Grocery": 100,
		"Petrol":  2199,
		"Movie":   956,
	}
	x := 956
	fmt.Println("")
	for key, value := range empBill {
		if value == x {
			fmt.Printf("key %d has value %s\n", value, key)
		}
	}
}

// Count words in a sentence
func SentenceWordCount() {
	line := "Quick brown fox jumps over the lazy dog."
	x := make(map[string]int)
	fmt.Println("\nCount words in a sentence")
	new := strings.Fields(line)
	for _, words := range new {
		words := strings.ToLower(words)
		x[words]++
	}
	for key, value := range x {
		fmt.Printf("Key is %s and the value is %d\n", key, value)
	}
}

// Find duplicate in a slice
func DuplicateSlice() {
	numbers := []int{10, 20, 30, 40, 50, 10, 60, 30}
	dup := make(map[int]int)

	fmt.Println("\nFinding duplicate in a slice")
	for _, num := range numbers {
		dup[num]++
	}
	for key, value := range dup {
		fmt.Printf("Key is %d and it appears %d time\n", key, value)
	}
}

// map of maps
func MapOfMaps() {
	outerMap := map[string]map[string]int{
		"John": {
			"Math":    85,
			"Science": 92,
		},
		"Jane": {
			"Math":    90,
			"Science": 88,
		},
	}
	fmt.Println("\nMap of Maps")
	valueToPrint := outerMap["John"]["Math"]
	fmt.Println(valueToPrint)
	fmt.Println("")
	outerMap["Jane"]["Bio"] = 80
	for key, value := range outerMap {
		fmt.Printf("The key of map of maps is %s\n", key)
		for innerkey, innervalue := range value {
			fmt.Printf("The inner key is %s and the inner value is %d\n", innerkey, innervalue)
		}
	}
}

// Flatten a Map of Slices
func FlattenMap() {
	fmt.Println("\nFlatten a Map of Slices")
	myMap := map[string][]int{
		"group1": {1, 2, 3},
		"group2": {4, 5},
		"group3": {6, 7, 8, 9},
	}
	x := []int{}
	for _, value := range myMap {
		fmt.Println(value)
		for _, v := range value {
			x = append(x, v)
		}
	}
	fmt.Println("new slice is ", x)
}

// Custom Sorting by Value
func SortingMap() {
	empBill := map[string]int{
		"Grocery": 100,
		"Petrol":  2199,
		"Movie":   956,
	}

	fmt.Println("\nSorting Map")
	keys := make([]string, 0, len(empBill))
	for key := range empBill {
		keys = append(keys, key)
	}
	for i := 0; i < len(empBill); i++ {
		for j := i + 1; j < len(empBill); j++ {
			if empBill[keys[i]] > empBill[keys[j]] {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}
	for _, key := range keys {
		fmt.Printf("%s:%d\n", key, empBill[key])
	}
}

// Convert Map to JSON
func MapJson() {
	empBill := map[string]int{
		"Grocery": 100,
		"Petrol":  2199,
		"Movie":   956,
	}
	fmt.Println("\nMap To Json")
	emp, err := json.Marshal(empBill)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(emp))
}

// Json Parsing
func JsonParsing() {
	jsonStr := `{
		"employee": {
			"name": "John Doe",
			"age": 30,
			"address": {
				"city": "New York",
				"zipcode": "10001"
			},
			"department": {
				"name": "Engineering",
				"id": 101
			}
		}
	}`

	fmt.Println("\nNested Map of Json parsing")
	var empData map[string]map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &empData)
	if err != nil {
		log.Fatal(err)
	}
	worker := empData["employee"]
	// Accessing values
	address := worker["address"].(map[string]interface{})
	fmt.Println("city: ", address["city"])
	fmt.Println("Zipcode", address["zipcode"])

}
