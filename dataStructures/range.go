package main

import "fmt"

func main() {

    nums := []int{2, 3, 4} //decalre
    sum := 0
    for _, num := range nums { //similar to reducer function, iterates over all in array
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums { // can also get index while ranging
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"} // can iterate over maps (nice feature)
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" { //range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself.
        fmt.Println(i, c)
    }
}