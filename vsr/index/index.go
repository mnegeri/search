package index 

import (
    "fmt"
    "io/ioutil"
    //"github.com/mnegeri/search/document"
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
