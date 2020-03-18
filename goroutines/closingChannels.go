package main

import "fmt"

// Closing a channel indicates that no more values will be sent on it.

func main() {
    jobs := make(chan int, 5) // 5 elements buffer
    done := make(chan bool)  //We await the worker using the synchronization approach we saw earlier.

    go func() { //async func
        for {
            j, more := <-jobs //reciever
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true // to await until read all "done"
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ { //send j into jobs
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs) //When we have no more jobs for the worker we’ll close the jobs channel.
    fmt.Println("sent all jobs")

    <-done // We await the worker using the synchronization approach we saw earlier.
}