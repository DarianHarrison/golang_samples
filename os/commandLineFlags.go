package main

import (
    "flag"
    "fmt"
)

//BUILD
//go build commandLineFlags.go

//IN CASE OF WINDOWS, else just run ./commandLineFlags
//run in terminal examples
//commandLineFlags.exe -word=opt -numb=7 -fork -svar=flag
//commandLineFlags.exe -word=opt
//commandLineFlags.exe -word=opt a1 a2 a3
//commandLineFlags.exe -word=opt a1 a2 a3 -numb=7
//commandLineFlags.exe -h
//commandLineFlags.exe -wat

func main() {

    wordPtr := flag.String("word", "foo", "a string") // flags based on words, default "foo", Use -h or --help flags "a string"

    numbPtr := flag.Int("numb", 42, "an int") // numb and fork flags, using a similar approach to the word flag.
    boolPtr := flag.Bool("fork", false, "a bool")//fork with default var false

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var") // default "bar", It’s also possible to declare an option that uses an existing var declared elsewhere in the program.

    flag.Parse() //flag.Parse() to execute the command-line parsing.

    //Here we’ll just dump out the parsed options and any trailing positional arguments.
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args()) //Trailing positional arguments
}