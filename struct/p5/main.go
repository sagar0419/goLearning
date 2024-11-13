package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func jsonEncode() {
	user := User{
		ID:        1,
		FirstName: "Sagar",
		LastName:  "Parmar",
		Email:     "sagar.rajput27@live.com",
	}

	usr, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(usr))
}

func jsonDecode() {
	usrDecode := []byte(`{
		"id":        2,
		"firstName": "Acer",
		"lastName":  "laptop",
		"email":     "acer.laptop@gmail.com"
	}`)

	var user User
	if err := json.Unmarshal(usrDecode, &user); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("User id: %d\nName: %s %s\nEmail: %s\n", user.ID, user.FirstName, user.LastName, user.Email)
}

func main() {
	jsonEncode()
	jsonDecode()
}
