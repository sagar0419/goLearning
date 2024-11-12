package json

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Employee struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Department string `json:"dept"`
}

var newEmployee = `{
					 "name":"Alice",
					 "age":30,
					 "dept":"IT"
				   }`

func EncodeJson() {

	employee := Employee{
		Name:       "Alice",
		Age:        30,
		Department: "IT",
	}
	emp, err := json.Marshal(employee)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(emp))
}

func DecodeJson() {

	var emp Employee

	if err := json.Unmarshal([]byte(newEmployee), &emp); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Name: %s, Age: %d, Department: %s\n", emp.Name, emp.Age, emp.Department)
}

// Reading JSON file
func ReadJson(jsonFile string) {

	file, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("\nThe content of the JSON file is below.\n\n%s\n", file)
}

// Writing to JSON file
func WriteJson(jsonFile string) {
	file, err := os.OpenFile(jsonFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// adding new line at the end of existing json file before appending data
	line := "\n" + newEmployee

	// Writing data
	_, err = file.WriteString(line)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("\nThe line %s is added in %s file\n", newEmployee, jsonFile)
}
