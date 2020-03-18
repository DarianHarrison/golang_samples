package main

import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)
 
// For more complex state we can use a mutex to safely access data across multiple goroutines.

func main() {

    var state = make(map[int]int) //state will be a map

    var mutex = &sync.Mutex{} //mutex will synchronize access to state

    var readOps uint64 //keep track of how many read and write operations
    var writeOps uint64

    for r := 0; r < 100; r++ { //100 goroutines to execute repeated reads against the state
        go func() {
            total := 0
            for { //For each read

                key := rand.Intn(5) //  we pick a key to access,
                mutex.Lock() // Lock() the mutex to ensure exclusive access to the state
                total += state[key] //  read the value at the chosen key, update state
                mutex.Unlock() // unlock
                atomic.AddUint64(&readOps, 1) //increment readops count

                time.Sleep(time.Millisecond) // Wait a bit between reads.
            }
        }()
    }

    for w := 0; w < 10; w++ { //start 10 goroutines to simulate writes
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second) //Let the 10 goroutines work on the state and mutex for a second.

    //Take and report final operation counts
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)

    mutex.Lock() //With a final lock of state
    fmt.Println("state:", state) //, show how it ended up.
    mutex.Unlock()
    
} //Running the program shows that we executed about 90,000 total operations against our mutex-synchronized state.