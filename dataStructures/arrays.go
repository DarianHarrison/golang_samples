package main

import "fmt"

func main() {

    var a [5]int //declare array of 5 ints
    fmt.Println("emp:", a)

    a[4] = 100 //set a value
    fmt.Println("set:", a)
    fmt.Println("get:", a[4]) //get a value

    fmt.Println("len:", len(a)) //length

    b := [5]int{1, 2, 3, 4, 5} //set multiple
    fmt.Println("dcl:", b)

    var twoD [2][3]int //multidim arrays
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}