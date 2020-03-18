package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) { //worker which we’ll run several concurrent instances. 
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {

	//In order to use our pool of workers we need to send them work and collect their results. We make 2 channels for this.
    const numJobs = 5
    jobs := make(chan int, numJobs) // buffer of 5
    results := make(chan int, numJobs) // buffer of 5

    for w := 1; w <= 3; w++ { //This starts up 3 workers, initially blocked because there are no jobs yet.
        go worker(w, jobs, results)
    }

    for j := 1; j <= numJobs; j++ { // send 5 jobs and then close
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= numJobs; a++ { //collect all the results of the work. This also ensures that the worker goroutines have finished.
        <-results
    }
}