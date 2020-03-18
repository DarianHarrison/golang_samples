package main

import (
    "fmt"
    "io/ioutil"
    "os/exec"
)

// if possible, run this on linux or mac, because it uses bash example


func main() {

    //simple command that takes no arguments or input and just prints something to stdout
    dateCmd := exec.Command("date") //The exec.Command helper creates an object to represent this external process.

    //.Output is another helper that handles the common case of running a command, waiting for it to finish, and collecting its output. If there were no errors, dateOut will hold bytes with the date info.
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))

    grepCmd := exec.Command("grep", "hello") //case where we pipe data to the external process on its stdin and collect the results from its stdout.

    
    grepIn, _ := grepCmd.StdinPipe()//Here we explicitly grab input/output pipes,
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start() // start the process,
    grepIn.Write([]byte("hello grep\ngoodbye grep")) // write some input to it,
    grepIn.Close()
    grepBytes, _ := ioutil.ReadAll(grepOut) // read the resulting output,
    grepCmd.Wait() // and finally wait for the process to exit.

    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes)) // print bytes from stdOut

    //when spawning commands we need to provide an explicitly delineated command and argument array, vs. being able to just pass in one command-line string.
    lsCmd := exec.Command("bash", "-c", "ls -a -l -h") //If you want to spawn a full command with a string, you can use bash’s -c option:
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}