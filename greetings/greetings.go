package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("hello %v", name)
	return message
}
