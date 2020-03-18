package main

import "fmt"

//methods are functions on objects

type rect struct {
    width, height int
}

func (r *rect) area() int { //takes in mem address
    return r.width * r.height
}

func (r rect) perim() int { //takes in the object here we call it a "struct"
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r // get mem address at any given time
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}