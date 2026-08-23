# Can’t seem to figure this out: all goroutines are asleep - deadlock!
https://forum.golangbridge.org/t/cant-seem-to-figure-this-out-all-goroutines-are-asleep-deadlock/4455

robert-mcdermott, Jan 2017, 
> I’m new to Go and have been struggling with the below code; I’ve tried many different things but I can’t seem to get it to run without getting a “fatal error: all goroutines are asleep - deadlock!” error. What am I missing?
``` Go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	messages := make(chan string)
	for x := 1; x <= 5; x++ {
		wg.Add(1)
		go sayhello(x, wg, &messages)
	}

	for msg := range messages {
		fmt.Println(msg)
	}
	wg.Wait()
	close(messages)
}

func sayhello(count int, wg *sync.WaitGroup, messages *chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(1000))
	*messages <- fmt.Sprintf("hello: %d", count)
}
```
Suneil Patel,
> Here is a modified version that worked for me. Comment out the line with close(messages) and it will deadlock. Using for range on a channel is the cause of the deadlock as it’ll never exit that loop.
``` Go
package main

import (
	"fmt"
	"sync"
	"time"
	"log"
)

func main() {
	wg := new(sync.WaitGroup)
	messages := make(chan string)
	for x := 1; x <= 5; x++ {
		wg.Add(1)
		go sayhello(x, wg, &messages)
	}

	go func(wg *sync.WaitGroup, messages chan string) {
		log.Println("waiting")
		wg.Wait()
		log.Println("done waiting")
		close(messages)
	}(wg, messages)

	for msg := range messages {
		fmt.Println(msg)
	}
}

func sayhello(count int, wg *sync.WaitGroup, messages *chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(1000 * count))
	*messages <- fmt.Sprintf("hello: %d", count)
	log.Println("sent message ", count)
}
```
christophberger,
> To explain the background of the problem: The range operator reads from the channel until the channel is closed. So in the original code, the for-range loop keeps waiting for more input from the channel, and wg.Wait() is never reached. Suneil’s solution separates reading the channel from waiting for the workers to finish.

Oleg,
``` Go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	messages := make(chan string, 5)
	for x := 1; x <= 5; x++ {
		wg.Add(1)
		go sayhello(x, wg, &messages)
	}
	wg.Wait()
	close(messages)
	for msg := range messages {
		fmt.Println(msg)
	}
}

func sayhello(count int, wg *sync.WaitGroup, messages *chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(1000))
	*messages <- fmt.Sprintf("hello: %d", count)
}
```
> You have to close the channel to notify the for loops to finish.--https://go.dev/play/p/nU2Rq6PJdO
``` Go
package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {
	message := make(chan string)
	go sayhello(message)
	for x := int64(1); x <= 5; x++ {
		m := "hello : " + strconv.FormatInt(x, 10)
		message <- m
	}
	wg.Wait()
}
func DisplayMsg(m string) {
	defer wg.Done()
	fmt.Printf("hello: %s\n", m)
	return
}
func sayhello(message chan string) {
	for {
		m, more := <-message
		wg.Add(1)
		go DisplayMsg(m)
		if more == false {
			close(message)
			return
		}
	}
}
```