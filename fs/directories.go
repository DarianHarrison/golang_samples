package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
)

// working with directories in the file system

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    err := os.Mkdir("subdir", 0755) // Create a new sub-directory in the current working directory.
    check(err)

    defer os.RemoveAll("subdir") // When creating temporary directories, it’s good practice to defer their removal. os.RemoveAll will delete a whole directory tree (similarly to rm -rf).

    createEmptyFile := func(name string) { // Helper function to create a new empty file.
        d := []byte("")
        check(ioutil.WriteFile(name, d, 0644))
    }

    createEmptyFile("subdir/file1") //call 

    err = os.MkdirAll("subdir/parent/child", 0755) //mkdir with permissions
    check(err)

    createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

    c, err := ioutil.ReadDir("subdir/parent") // ReadDir lists directory contents, returning a slice of os.FileInfo objects.
    check(err)

    fmt.Println("Listing subdir/parent") 
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    err = os.Chdir("subdir/parent/child") // Chdir lets us change the current working directory, similarly to cd.
    check(err)

    c, err = ioutil.ReadDir(".") //Now we’ll see the contents of subdir/parent/child when listing the current directory.
    check(err)

    fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    err = os.Chdir("../../..") //cd back to where we started.
    check(err)

    fmt.Println("Visiting subdir") //We can also visit a directory recursively, including all its sub-directories. Walk accepts a callback function to handle every file or directory visited
    err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error { //visit is called for every file or directory found recursively by filepath.Walk.
    if err != nil {
        return err
    }
    fmt.Println(" ", p, info.IsDir())
    return nil
}