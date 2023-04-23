package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("in the init func %v", time.Now().UnixNano())
	fmt.Println()
	fmt.Println(time.Now())
	fmt.Println()
}

func randomFormat() string {
	formats := []string{
		"Hallo, %v. Great to see you",
		"Wilkomen, %v.",
		"Fifi hake, %v!!",
	}
	return formats[rand.Intn(len(formats))]
}
