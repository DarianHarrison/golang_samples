package main

import "fmt"

func main() {

    queue := make(chan string, 2) //channel, type string, buffer 2
    queue <- "one"
    queue <- "two"
    close(queue) // channel is now closed but we still are left with the queue buffer

    for elem := range queue { // note this empties the queue
        fmt.Println(elem)
    }

}