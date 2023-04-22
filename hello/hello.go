package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Gladiola")
	fmt.Println(message)
}
