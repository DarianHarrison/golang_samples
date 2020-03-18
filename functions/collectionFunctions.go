package main

import (
    "fmt"
    "strings"
)

//map, filter

func Index(vs []string, t string) int { // Index returns the first index of the target string t, or -1 if no match is found.
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

func Include(vs []string, t string) bool { // Include returns true if the target string t is in the slice.
    return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

func All(vs []string, f func(string) bool) bool { //Any returns true if one of the strings in the slice satisfies the predicate f. has prfix p
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

func Filter(vs []string, f func(string) bool) []string { //filter returns new slice with, all that have prefix p
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

func Map(vs []string, f func(string) string) []string { // Map returns a new slice w' results of applying the function to each elem. in original slice
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

func main() {

    var strs = []string{"peach", "apple", "pear", "plum"}

    fmt.Println(Index(strs, "pear"))

    fmt.Println(Include(strs, "grape"))

    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
    }))

    fmt.Println(Map(strs, strings.ToUpper))

}