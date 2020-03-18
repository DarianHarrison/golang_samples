package main

import (
    "fmt"
    "time"
)

// Here’s an example of using a blocking receive to wait for a goroutine to finish.

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() { 

    done := make(chan bool, 1) //The done channel will be used to notify another goroutine that this function’s work is done. 
    go worker(done) //Start a worker goroutine, giving it the channel to notify on

    <-done // Block until we receive a notification from the worker on the channel.
}