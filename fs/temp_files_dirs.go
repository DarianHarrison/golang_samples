package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
)

//Throughout program execution, we often want to create data that isn’t needed after the program exits.
//Temporary files and directories are useful for this purpose since they don’t pollute the file system over time.

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    f, err := ioutil.TempFile("", "sample") // easiest way to create a temporary file is by calling ioutil.TempFile.
    check(err)

    fmt.Println("Temp file name:", f.Name())

    defer os.Remove(f.Name()) // Clean up the file after we’re done. The OS is likely to clean up temporary files by itself after some time, but it’s good practice to do this explicitly.

    _, err = f.Write([]byte{1, 2, 3, 4}) // We can write some data to the file.
    check(err)

    dname, err := ioutil.TempDir("", "sampledir") // If we intend to write many temporary files, we may prefer to create a temporary directory. ioutil.
    fmt.Println("Temp dir name:", dname)

    defer os.RemoveAll(dname)

    fname := filepath.Join(dname, "file1") // Now we can synthesize temporary file names by prefixing them with our temporary directory.
    err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
    check(err)
}