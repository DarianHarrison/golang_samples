package main

import (
    "fmt"
    "sync"
    "time"
)

// To wait for multiple goroutines to finish, we can use a wait group.

func worker(id int, wg *sync.WaitGroup) { //This is the function we’ll run in every goroutine. Note that a WaitGroup must be passed to functions by pointer.

    defer wg.Done() //On return, notify the WaitGroup that we’re done.

    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second) //Sleep to simulate an expensive task.
    fmt.Printf("Worker %d done\n", id)
}

func main() {

    var wg sync.WaitGroup //This WaitGroup is used to wait for all the goroutines launched here to finish.

    for i := 1; i <= 5; i++ { //Launch several goroutines and increment the WaitGroup counter for each.
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait() //Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
}