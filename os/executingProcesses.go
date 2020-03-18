package main

import (
    "os"
    "os/exec"
    "syscall"
)

//best on linux or mac

//Sometimes we just want to completely replace the current Go process with another (perhaps non-Go) one.

func main() {

    // Go requires an absolute path to the binary we want to execute, so we’ll use exec.LookPath to find it (probably /bin/ls).
    // Exec requires arguments in slice form (as apposed to one big string).
    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

    args := []string{"ls", "-a", "-l", "-h"} //Exec requires arguments in slice form (as apposed to one big string). first argument should be the program name

    //Exec also needs a set of environment variables to use. Here we just provide our current environment.
    env := os.Environ()

    execErr := syscall.Exec(binary, args, env) //Here’s the actual syscall.Exec call.
    //If this call is successful, the execution of our process will end here and be replaced by the /bin/ls -a -l -h process.
    if execErr != nil {// If there is an error we’ll get a return value.
        panic(execErr)
    }
}