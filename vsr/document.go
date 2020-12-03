package vsr 

import (
   "bufio"
    "os"
    "fmt"
)

type Document struct {
    Name string 
    StopWords map[string]bool
}


func (doc *Document) LoadStopWords() {
    file, err := os.Open("stop_words.txt")
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        word := scanner.Text()
        fmt.Println(word)
        doc.StopWords[word] = true
    }
    file.Close()
}

func (doc *Document) Test() {
    fmt.Println(doc.Name)
}
