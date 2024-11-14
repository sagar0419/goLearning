package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Team struct {
	Name    string    `json:"name"`
	Players []Players `json:"players"`
}

type Players struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Position  string `json:"position"`
}

func main() {
	team := Team{
		Name: "BlackBull",
		Players: []Players{
			{FirstName: "James", LastName: "Smith", Position: "First"},
			{FirstName: "Johnny", LastName: "Walker", Position: "Second"},
			{FirstName: "Black", LastName: "Dog", Position: "Third"},
			{FirstName: "Blue", LastName: "Label", Position: "Fourth"},
			{FirstName: "Old", LastName: "Monk", Position: "Fifth"},
			{FirstName: "Blenders", LastName: "Pride", Position: "Sixth"},
		},
	}
	// Encode
	fmt.Println("Encoding Slice Struct")
	playerJson, err := json.MarshalIndent(team, "", " ")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(playerJson))

	// Decode
	fmt.Println("\n Decoding slice struct")

	var decodePlayerJson Team
	err = json.Unmarshal(playerJson, &decodePlayerJson)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Decode values are:- \n%+v", decodePlayerJson)
}
