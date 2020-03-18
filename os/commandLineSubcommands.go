package main

import (
    "flag"
    "fmt"
    "os"
)

// Some command-line tools, like the go tool or git have many subcommands, each with its own set of flags. For example, go build and go get are two different subcommands of the go tool.

//go build commandLineSubcommands.go 

//./commandLineSubcommands.exe foo -enable -name=joe a1 a2
//./commandLineSubcommands.exe bar -level 8 a1
//./commandLineSubcommands.exe bar -enable a1

func main() {

    
    fooCmd := flag.NewFlagSet("foo", flag.ExitOnError) //We declare a subcommand using the NewFlagSet function, 
    fooEnable := fooCmd.Bool("enable", false, "enable") // flag
    fooName := fooCmd.String("name", "", "name") // subcommand.

    barCmd := flag.NewFlagSet("bar", flag.ExitOnError) //flag
    barLevel := barCmd.Int("level", 0, "level") //subcommand

    if len(os.Args) < 2 { //The subcommand is expected as the first argument to the program.
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }

    switch os.Args[1] { //Check which subcommand is invoked.

    // For every subcommand, we parse its own flags and have access to trailing positional arguments.
    case "foo":
        fooCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'foo'")
        fmt.Println("  enable:", *fooEnable)
        fmt.Println("  name:", *fooName)
        fmt.Println("  tail:", fooCmd.Args())
    case "bar":
        barCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'bar'")
        fmt.Println("  level:", *barLevel)
        fmt.Println("  tail:", barCmd.Args())
    default:
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }
}