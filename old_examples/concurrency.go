//gorutine concurrency with synchronization.
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	for i:=0; i < 3; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}	
	wg.Done()
}

func main() {
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("there")
	wg.Wait()
}