package main

import (
	"example/go_module"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	names := []string{"Alice", "Bob"}

	msg, err := go_module.HelloList(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
