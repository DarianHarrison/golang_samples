package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

//  This channel-based approach aligns with Go’s ideas of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine.
// In order to read or write that state, other goroutines will send messages to the owning goroutine and receive corresponding replies.
// readOp and writeOp structs encapsulate those requests and a way for the owning goroutine to respond.
type readOp struct {
    key  int
    resp chan int
}

type writeOp struct {
    key  int
    val  int
    resp chan bool
}

func main() {

    var readOps uint64 // count ops
    var writeOps uint64

    reads := make(chan readOp) // channels to issue read write reqs, takes in type struct (object).
    writes := make(chan writeOp)

    go func() { // In this example our state will be owned by a single goroutine. This will guarantee that the data is never corrupted with concurrent access.
        var state = make(map[int]int)
        for {  //
            select { // This goroutine repeatedly selects on the reads and writes channels, responding to requests as they arrive.
            case read := <-reads: //A response is executed by first performing the requested operation
                read.resp <- state[read.key]
            case write := <-writes: // and then sending a value on the response channel resp to indicate success
                state[write.key] = write.val 
                write.resp <- true
            }
        }
    }()

    for r := 0; r < 100; r++ { // This starts 100 goroutines to issue reads to the state-owning goroutine via the reads channel
        go func() { // Each read requires constructing a readOp,  
            for {
                read := readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int)}
                reads <- read // sending it over the reads channel, 
                <-read.resp // and the receiving the result over the provided resp channel.
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    for w := 0; w < 10; w++ { // We start 10 writes as well, using a similar approach.
        go func() {
            for {
                write := writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second) //Let the goroutines work for a second.

    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}