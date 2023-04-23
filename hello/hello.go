package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")

	names := []string{"Arafael", "Beziquiel", "Curumden"}
	messages, err1 := greetings.Hellos(names)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(messages)

	message, err := greetings.Hello("Gladiolaaaa")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
