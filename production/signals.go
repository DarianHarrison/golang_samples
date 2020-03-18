package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

// intelligently handle Unix signals.
//For example, we might want a server to gracefully shutdown when it receives a SIGTERM, or a command-line tool to stop processing input if it receives a SIGINT.

func main() {// signal notification works by sending os.Signal values on a channel. 

    sigs := make(chan os.Signal, 1) // We’ll create a channel to receive these notifications 
    done := make(chan bool, 1) // we’ll also make one to notify us when the program can exit.

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // signal.Notify registers the given channel to receive notifications

    go func() { //This goroutine executes a blocking receive for signals. When it gets one it’ll print it out and then notify the program that it can finish.
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    fmt.Println("awaiting signal")
    <-done //The program will wait here until it gets the expected signal (as indicated by the goroutine above sending a value on done) and then exit.
    fmt.Println("exiting")
}