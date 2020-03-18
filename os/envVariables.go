package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {

    os.Setenv("FOO", "1") //To set a key/value pair, use os.Setenv.
    fmt.Println("FOO:", os.Getenv("FOO")) //To get a value for a key, use os.Getenv. 
    fmt.Println("BAR:", os.Getenv("BAR"))

    fmt.Println()
    for _, e := range os.Environ() {//Use os.Environ to list all key/value pairs in the environment.
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
}