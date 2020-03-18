// We use Go Channels to connect concurrent goroutines, in order to send and receive values between them, using the channel operator <-
//To send the value over a channel, you would do c <- v. To receive a value from a channel, you would do var := <-c.

package main

import "fmt"

func foo(c chan int, someValue int) {
    c <- someValue * 5
}

func main() {
    fooVal := make(chan int)
    go foo(fooVal, 5)
    go foo(fooVal, 3)
    v1 := <-fooVal
    v2 := <-fooVal
    fmt.Println(v1, v2)
}