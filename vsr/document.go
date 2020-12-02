package vsr 

import (
   "bufio"
    "os"
    "fmt"
)

type Document struct {
    Name string 
    stopWords map[string]bool
}


func (doc Document) LoadStopWords() {
    stopWords := make(map[string]bool)
    file, err := os.Open("stop_words.txt")
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        stopWords[scanner.Text()] = true
    }
    file.Close()
}

func (doc Document) Test() {
    fmt.Println(doc.Name)
}
