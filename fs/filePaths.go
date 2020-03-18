package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

//The filepath package provides functions to parse and construct file paths in a way that is portable between operating systems; dir/file on Linux vs. dir\file on Windows, for example.

func main() {

    p := filepath.Join("dir1", "dir2", "filename") //Join takes any number of arguments and constructs a hierarchical path from them.
    fmt.Println("p:", p)

    fmt.Println(filepath.Join("dir1//", "filename")) //You should always use Join instead of concatenating strings
    fmt.Println(filepath.Join("dir1/../dir1", "filename"))

    fmt.Println("Dir(p):", filepath.Dir(p)) //Dir and Base can be used to split a path to the directory and the file.
    fmt.Println("Base(p):", filepath.Base(p))

    fmt.Println(filepath.IsAbs("dir/file")) //We can check whether a path is absolute.
    fmt.Println(filepath.IsAbs("/dir/file"))

    filename := "config.json"

    ext := filepath.Ext(filename) //We can split the extension out of such names with Ext.
    fmt.Println(ext)

    fmt.Println(strings.TrimSuffix(filename, ext))//file’s name with the extension removed, use strings.TrimSuffix.

    rel, err := filepath.Rel("a/b", "a/b/t/file")//Rel finds a relative path between a base and a target.
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)

    rel, err = filepath.Rel("a/b", "a/c/t/file") //
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
}