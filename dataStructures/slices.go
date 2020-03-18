package main

import "fmt"

func main() {
	//create and set arary
    s := make([]string, 3)
    fmt.Println("emp:", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2]) //get test

    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    c := make([]string, len(s)) //make new array
    copy(c, s) //copy values from array
    fmt.Println("cpy:", c)

    
    l := s[2:5] //2 limit slice
    fmt.Println("sl1:", l) 

    l = s[:5] //all till five
    fmt.Println("sl2:", l)

    l = s[2:]//index to till end
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"} 
    fmt.Println("dcl:", t)//create new array t with 3 elems

    twoD := make([][]int, 3) //two dim array
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}