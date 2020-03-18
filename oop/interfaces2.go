package main

import (
    "fmt"
)

type Animal interface {
    Speak() string
}

type Dog struct {

}

func (d Dog) Speak() string { //takes in type dog, funct name (args), return type string
    return "Woof!"
}

type Cat struct {

}

func (c Cat) Speak() string {
    return "Meow!"
}

type Llama struct {

}

func (l Llama) Speak() string {
    return "?????"
}

type JavaProgrammer struct {

}

func (j JavaProgrammer) Speak() string {
    return "Design patterns!"
}

func main() {
    animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}} //create array of 4 Animal interphaces
    for _, animal := range animals { // if return false _. do nothing, else return animal
        fmt.Println(animal.Speak()) // call methods
}