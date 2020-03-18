//The idea of the defer statement is to put off (defer) the execution of something until the surrounding function is done.
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()//to make sure it runs, will be evaluated immediately, but only runs after rest of say function is done.
	for i:=0; i < 3; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}	
}

func main() {
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("there")
	wg.Wait()
}