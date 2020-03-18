package main

import (
    "fmt"
    "os"
)

//cd to directory
//go build commandLineArgs.go
//commandLineArgs.exe a b c d


func main() {

    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}