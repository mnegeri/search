//Package vsr provides structures and methods for developing an 
//information retrieval system using a vector space model.
package vsr 

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

//Document represents methods that are implemented by various document types.
type Document interface {
    HashMapVector() *Vector
}

//FileDocument represents a file document.
type FileDocument struct {
    //FilePath must be the absolute path
    FilePath string 
    VectorLength float64
    //StopWords map[string]bool
    SimilarityScore float64
}


//HashMapVector loops through every term of a document and add terms that
//are not stop words to a vector and returns the vector.
func (doc *FileDocument) HashMapVector() *Vector {
    stopWords := LoadStopWords()
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
}

//QueryDocument document represents a Document constructed from a string query
type QueryDocument struct {
    Query string
    VectorLength float64
    SimilarityScore float64
}

//HashMapVector loops through every term of a query and add terms that
//are not stop words to a vector and returns the vector.
func (doc *QueryDocument) HashMapVector() *Vector {
    stopWords := LoadStopWords()
    vector := Vector{make(map[string]float64)}
    terms := strings.Fields(doc.Query) 
    for _, term := range terms {
        if !stopWords[term] {
            vector.Add(term)
        }
    }
    return &vector
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

