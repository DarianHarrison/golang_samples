package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

// There are a few other options for managing state. Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines

func main() {

    var ops uint64 // unsigned integer to represent our (always-positive) counter.

    var wg sync.WaitGroup // WaitGroup will help us wait for all goroutines to finish their work.

    for i := 0; i < 50; i++ { //50 goroutines that each increment the counter exactly 1000 times
        wg.Add(1)

        go func() {
            for c := 0; c < 1000; c++ {

                atomic.AddUint64(&ops, 1)
            }
            wg.Done()
        }()
    }

    wg.Wait() //wait for wg. to be done

    fmt.Println("ops:", ops) // It’s safe to access ops now because we know no other goroutine is writing to it.
}