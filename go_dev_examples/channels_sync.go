// https://stackoverflow.com/questions/59294918/golang-pass-struct-between-two-goroutines-concurrently
package main

import (
	"fmt"
	"sync"
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

	// Spin up the workers to find people.
	var finders sync.WaitGroup
	for i := 0; i < 5; i++ {
		finders.Add(1)
		go func() {
			defer finders.Done()
			FindPeople(people)
		}()
	}

	// Close chan people once all finders have finished.
	go func() {
		finders.Wait()
		close(people)
	}()

	// Read the list of people from the find operation.
	var writers sync.WaitGroup
	for i := 0; i < 3; i++ {
		writers.Add(1)
		go func() {
			defer writers.Done()
			for p := range people {
				WritePerson(p)
			}
		}()
	}

	// Wait for the writers to finish.
	writers.Wait()
	fmt.Println("Finished")
}
