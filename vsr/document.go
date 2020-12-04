package vsr 

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

type Document struct {
    //FilePath must be the absolute path
    FilePath string 
    StopWords map[string]bool
}

const stopwords = `a an and are as at be by for from has he in is it its of on that the to was were will with`


func (doc *Document) LoadStopWords() {
    /*file, err := os.Open("vsr/stop_words.txt")
    if err != nil {
        panic(err)
    } 
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        word := scanner.Text()
        fmt.Println(word)
        doc.StopWords[word] = true
    }
    file.Close() */
    words := strings.Fields(stopwords)
    for _, word := range words {
        doc.StopWords[word] = true
    }
}

func (doc *Document) HashMapVector() *Vector {
    file, err := os.Open(doc.FilePath)
    if err != nil {
        panic(err)
    }
    vector := Vector{make(map[string]int)}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        token := scanner.Text()
        if !doc.StopWords[token] {
            vector.Add(token)
        }
    }
    file.Close() 
    return &vector
}

func (doc *Document) Test() {
    fmt.Println(doc.FilePath)
}
