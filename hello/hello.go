package main

import (
	"fmt"
	"imtzz.com/greetings"
	"log"
	"rsc.io/quote"
)

//func callGreetings(name string) {
//	message := greetings.Hello(name)
//	fmt.Println(message)
//}

func callGreetings(name string) {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

func callNamesGreetings(names []string) {
	log.SetPrefix("callNamesGreetings:")
	log.SetFlags(0)

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}

func main() {
	fmt.Printf("Hello World!\n")

	fmt.Println(quote.Hello())

	//callGreetings("Gladys")
	//
	//callGreetings("")

	names := []string{"Gladys", "samantha", "Darrin"}
	callNamesGreetings(names)
}
