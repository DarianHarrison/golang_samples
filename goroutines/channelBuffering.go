package main

import "fmt"

// Buffered channels accept a limited number of values without a corresponding receiver for those values.
func main() {

    messages := make(chan string, 2) //buffering up to 2 values

    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}