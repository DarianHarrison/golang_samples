package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    d1 := []byte("hello\ngo\n") // dump a string (or just bytes) into a file
    err := ioutil.WriteFile("dat/dat1.txt", d1, 0644)
    check(err)

    f, err := os.Create("dat/dat2.txt") // For more granular writes, open a file for writing
    check(err)

    defer f.Close() // deferred file close

    d2 := []byte{115, 111, 109, 101, 10} // You can Write byte slices as you’d expect
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    n3, err := f.WriteString("writes\n") //A WriteString is also available.
    fmt.Printf("wrote %d bytes\n", n3)

    f.Sync() // Issue a Sync to flush writes to stable storage.

    w := bufio.NewWriter(f) // bufio provides buffered writers in addition to the buffered readers we saw earlier.
    n4, err := w.WriteString("buffered\n")
    fmt.Printf("wrote %d bytes\n", n4)

    w.Flush() //Flush to ensure all buffered operations have been applied to the underlying writer

}