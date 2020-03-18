package main

import (
    "fmt"
    "time"
)

// Timers are for when you want to do something once in the future

func main() {

    timer1 := time.NewTimer(2 * time.Second)

    <-timer1.C //The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer fired.
    fmt.Println("Timer 1 fired")

    timer2 := time.NewTimer(time.Second)

    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    stop2 := timer2.Stop()

    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

    time.Sleep(2 * time.Second)
}