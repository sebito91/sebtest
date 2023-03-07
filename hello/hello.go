package main

import (
	"fmt"
	"log"

	"github.com/sebito91/sebtest/greetings"
)

func main() {
	fmt.Println("vim-go")

	log.SetPrefix("greetings yo! ")
	log.SetFlags(0)

	msg, err := greetings.Hello("Bobski")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)

	msg, err = greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}
