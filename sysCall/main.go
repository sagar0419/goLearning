package main

import (
	"log"
	s "syscall"
)

func main() {
	log.Println("Getpid : ", s.Getpid())
	buffer := make([]byte, 1024)

	length, err := s.Getcwd(buffer)
	if err != nil {
		log.Fatal("Connot get current Working dorectory error occured: ", err)
	}
	log.Println("Current working directory: ", string(buffer[:length]))
}
