package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    dat, err := ioutil.ReadFile("dat/dat1.txt") // file’s entire contents into memory
    check(err)
    fmt.Print(string(dat))

    f, err := os.Open("dat/dat1.txt")
    check(err)

    b1 := make([]byte, 5) // Read some bytes from the beginning of the file
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

    o2, err := f.Seek(6, 0) // Seek to a known location in the file and Read from there
    check(err)
    b2 := make([]byte, 2) 
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: ", n2, o2)
    fmt.Printf("%v\n", string(b2[:n2]))

    o3, err := f.Seek(6, 0) //read using io package.
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    _, err = f.Seek(0, 0) // There is no built-in rewind, but Seek(0, 0) accomplishes this.
    check(err)

    r4 := bufio.NewReader(f) // bufio package implements a buffered reader
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    f.Close() //Close the file when you’re done (usually this would be scheduled immediately after Opening with defer).
}