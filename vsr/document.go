//Package vsr provides structures and methods for developing an 
//information retrieval system using a vector space model.
package vsr 

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

type Doc interface {
    HashMapVector() *Vector
}

//Document represents a file document.
type Document struct {
    //FilePath must be the absolute path
    FilePath string 
    StringText string
    VectorLength float64
    //StopWords map[string]bool
    SimilarityScore float64
}

const stopwords = `a an and are as at be by for from has he in is it its of on that the to was were will with`


//LoadStopWords creates a hash map of stop words.
func LoadStopWords() map[string]bool {
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
    stopWords := make(map[string]bool)
    for _, word := range words {
        stopWords[word] = true
    }
    return stopWords
}

//HashMapVector loops through every term of a document and add terms that
//are not stop words to a vector and returns the vector.
func (doc *Document) HashMapVector() *Vector {
    stopWords := LoadStopWords()
    if (doc.StringText == "") {
        file, err := os.Open(doc.FilePath)
        if err != nil {
            panic(err)
        }
        vector := Vector{make(map[string]float64)}
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            token := scanner.Text()
            strings.ToLower(token)
            if !stopWords[token] {
                vector.Add(token)
            }
        }
        file.Close() 
        return &vector
    } else {
        vector := Vector{make(map[string]float64)}
        terms := strings.Fields(doc.StringText) 
        for _, term := range terms {
            vector.Add(term)
        }
        return &vector
    }
}

func (doc *Document) Test() {
    fmt.Println(doc.FilePath)
}
