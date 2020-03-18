package main

import "os"
//we use it to fail fast on errors that shouldn’t occur during normal operation
func main() {

    panic("a problem") //prints "a problem" if error 

    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}