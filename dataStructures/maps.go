package main

import "fmt"

func main() { //in Go, Maps are Key-Value data structures, NOT map function

    m := make(map[string]int) //delcare map

    m["k1"] = 7 //set vals
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"] //get 7
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m)) //print len of map (nice feature)

    delete(m, "k2")
    fmt.Println("map:", m)

    _, prs := m["k2"] //The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _.
    fmt.Println("prs:", prs) 

    n := map[string]int{"foo": 1, "bar": 2} //declare and initialize a new map
    fmt.Println("map:", n)
}