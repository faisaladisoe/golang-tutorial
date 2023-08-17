package greetings

import "fmt"

func Hello(name string) string {
	greet := fmt.Sprintf("Hi, %s. Welcome!", name)
	return greet
}