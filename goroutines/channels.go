package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines. 

func main() {

    messages := make(chan string) //create a new channel of type string

    go func() { messages <- "ping" }() //send value into a channel 

    msg := <- messages //receive value sent through channel 
    fmt.Println(msg) //print out received channel 
}

//"ping" message is successfully passed from one goroutine to another via our channel