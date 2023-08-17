package main

import (
	"fmt"
	"math/rand"

	"rsc.io/quote"
)

func main() {
	fmt.Print(rand.Intn(100), "\n")
	fmt.Printf("string formatting: %s %t %d %f\n", "saya pejalan kaki", false, 10, 10.5)
	fmt.Printf("string formatting using v: %v %v %v\n", 12, "I am", true)
    fmt.Println(quote.Glass())
    fmt.Println(quote.Go())
    fmt.Println(quote.Hello())
    fmt.Println(quote.Opt())
}