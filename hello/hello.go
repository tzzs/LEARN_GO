package main

import (
	"fmt"
	"imtzz.com/greetings"
	"log"
	"rsc.io/quote"
)

//func callGreetins(name string) {
//	message := greetings.Hello(name)
//	fmt.Println(message)
//}

func callGreetins(name string) {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

func main() {
	fmt.Printf("Hello World!\n")

	fmt.Println(quote.Hello())

	callGreetins("Gladys")

	callGreetins("")
}
