package main

import (
    "fmt"
    "time"
)

//for controlling resource utilization and maintaining quality of service.

func main() {

    requests := make(chan int, 5) 
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(200 * time.Millisecond) // can only receive a value every 200 millis

    for req := range requests { // By blocking on a receive from the limiter channel before serving each request, we limit ourselves to 1 request every 200 milliseconds.
        <-limiter
        fmt.Println("request", req, time.Now())
    }

    burstyLimiter := make(chan time.Time, 3) //This burstyLimiter channel will allow bursts of up to 3 events

    for i := 0; i < 3; i++ { //Fill up the channel
        burstyLimiter <- time.Now()
    }

    go func() {//Every 200 milliseconds we’ll try to add a new value to burstyLimiter, up to its limit of 3
        for t := range time.Tick(200 * time.Millisecond) { 
            burstyLimiter <- t
        }
    }()

    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ { // Now simulate 5 more incoming requests.
        burstyRequests <- i
    }
    close(burstyRequests)

    for req := range burstyRequests { // the first batch of requests handled once every ~200 milliseconds
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}