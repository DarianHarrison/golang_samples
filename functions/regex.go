package main

import (
    "bytes"
    "fmt"
    "regexp"
)

//tests whether a pattern matches a string

var p = fmt.Println // since we'll print a lot

func main() {

    r, _ := regexp.Compile("p([a-z]+)ch")

    p(r.MatchString("peach")) //Here’s a match test

    p(r.FindString("peach punch")) //finds the match

    p(r.FindStringIndex("peach punch")) //returns the start and end indexes

    p(r.FindStringSubmatch("peach punch")) //of first ocurrence

    p(r.FindStringSubmatchIndex("peach punch"))

    p(r.FindAllString("peach punch pinch", -1)) //returns array of all

    p(r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1)) // array of indices

    p(r.FindAllString("peach punch pinch", 2)) // limit 2 finds

    p(r.Match([]byte("peach")))

    r = regexp.MustCompile("p([a-z]+)ch")
    p(r)

    p(r.ReplaceAllString("a peach", "<fruit>"))

    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    p(string(out))
}