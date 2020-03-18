package main

//Use os.Exit to immediately exit with a given status.

import (
    "fmt"
    "os"
)

func main() {

    defer fmt.Println("!") //defers will not be run when using os.Exit, so this fmt.Println will never be called.

    os.Exit(3)Exit with status 3.
}