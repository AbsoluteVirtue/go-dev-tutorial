// https://stackoverflow.com/questions/59294918/golang-pass-struct-between-two-goroutines-concurrently
package main

import (
	"fmt"
	"time"
)

type Person struct {
	index int
	name  string
	age   int
}

func FindPeople(people chan Person) {
	for i := 0; i < 5; i++ {
		p := Person{
			index: i,
			name:  "Shaun",
			age:   23,
		}
		fmt.Println("Found Person", p)
		people <- p
		time.Sleep(time.Millisecond)
	}
}

func WritePerson(p Person) {
	fmt.Println("Writing People to DB", p)
}

func main() {
	people := make(chan Person, 0)

	// Do whatever work to produce our list of people, then close it.
	go func() {
		FindPeople(people)
		close(people)
		fmt.Println("Find complete.")
	}()

	// Read the list of people from the find operation.
	for p := range people {
		WritePerson(p)
	}
	fmt.Println("Finished")
}
