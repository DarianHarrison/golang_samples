package main

import "fmt"

//  use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    select { //If a value is available on messages then select will take the <-messages case with that value. If not it will immediately take the default case.
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    msg := "hi"
    select { //Here msg cannot be sent to the messages channel, because the channel has no buffer and there is no receiver.
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    select { //use multiple cases above the default clause to implement a multi-way non-blocking select.
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}