package main

import (
    "fmt"
    "os"
)

//Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err) //note the panic
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) { //check for errors even in a deferred function.
    fmt.Println("closing")
    err := f.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

func main() { //Suppose we wanted to create a file, write to it, and then close when we’re done. 

    f := createFile("defer.txt")
    defer closeFile(f) // execute this last
    writeFile(f)
}

