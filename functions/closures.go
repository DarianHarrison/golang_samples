package main

import "fmt"

//closure is an instance of anon function that can have local state
//this can be good for sampling and RL

func intSeq() func() int { //outer function returns inner function (this case takes no input)
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

    nextInt := intSeq() //here we are returning an instance of the function func () 

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq() // declare another instance of funct () proove state is unique to previous function
    fmt.Println(newInts())
    fmt.Println(newInts())
}