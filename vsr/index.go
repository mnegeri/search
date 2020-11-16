package vsr

import (
    "fmt"
    "io/ioutil"
)

type InvertedIndex struct {
    stem bool
    html bool
    dir string
}

func (i InvertedIndex) Process() {
    files, err := ioutil.ReadDir(dir)
    if (err != nill) {
        log.Fatal(err)
    }
    for _, file := range files {
        fmt.Println(file.Name())
    }
}
