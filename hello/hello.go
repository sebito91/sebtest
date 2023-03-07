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

	msg, err := greetings.Hello("Steve")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}
