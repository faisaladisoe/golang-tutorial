package main

import (
	"fmt"
	"example/greetings"
)

func main() {
	message := greetings.Hello("Crystal")
	fmt.Println(message)
}