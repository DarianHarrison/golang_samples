package main

import "fmt"

func zeroval(ival int) {
    ival = 0
    //1 = 0
}

func zeroptr(iptr *int) {//*int parameter, meaning that it takes an int pointer.
    *iptr = 0 //pointer of variable with name 1
}

func main() {
    i := 1
    fmt.Println("initial:", i) //print above declaration

    zeroval(i)
    fmt.Println("zeroval:", i) //1 = 0, i is still 1

    zeroptr(&i)//memory address of i will now equal 0
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i) //print address of var with name i

    //example 2
    var a = 5
    var p = &a //p holds varibale's a as mem address
    fmt.Println(p)
    fmt.Println(*p) //get value at this address

    *p = 3
    a = 4 //change value of a
    fmt.Println(p)
    fmt.Println(*p) //get value at this address

}