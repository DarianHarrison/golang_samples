//The idea of panic is to halt the program and start to panic. This will stop the running function and run all of the deferred functions from the panicking function. The recover function lets you recover from a panicking goroutine. To recover a panicking goroutine, you would need to use recover from within one of the deferred functions of that goroutine.
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup:", r)
	}
	wg.Done()
}

func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
		if i == 2 {
			panic("Oh dear... a 2")
		}
	}	
}

func main() {
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("there")
	wg.Wait()
}