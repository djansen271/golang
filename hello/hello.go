package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	message, err := greetings.Hello("Gladiolaaaa")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
