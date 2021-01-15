//Package vsr provides structures and methods for developing an 
//information retrieval system using a vector space model.
package vsr 

import (
    "bufio"
    "os"
    "strings"
    "path/filepath"
    "fmt"
)

//Document represents methods that are implemented by various document types (query doc, text file doc, html doc).
type Document interface {
    HashMapVector() *Vector
    GetVectorLength() float64
    SetVectorLength(val float64) 
    GetSimilarityScore() float64
    SetSimilarityScore(val float64)
    GetName() string
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
func (doc FileDocument) HashMapVector() *Vector {
    stopWords := LoadStopWords()
    file, err := os.Open(doc.FilePath)
    if err != nil {
        panic(err)
    }
    vector := Vector{make(map[string]float64)}
    scanner := bufio.NewScanner(file)
    //loop through each term in the file
    for scanner.Scan() {
        term := scanner.Text()
        strings.ToLower(term)
        //if the term is not a stop word, add it to the vector
        if !stopWords[term] {
            vector.Add(term)
        }
    }
    file.Close() 
    return &vector
}

//GetVectorLength returns the current value of the vector lenght of a document
func (doc FileDocument) GetVectorLength() float64 {
    return doc.VectorLength
}

//SetVectorLength sets the value of the vector length of this document to val
func (doc *FileDocument) SetVectorLength(val float64) {
    doc.VectorLength = val
}

//GetSimilarityScore returns the current value of the similarity score  
func (doc FileDocument) GetSimilarityScore() float64 {
    return doc.SimilarityScore
}

//SetSimilarityScore sets the value of the similarity score for this document to val
func (doc *FileDocument) SetSimilarityScore(val float64) {
    doc.SimilarityScore = val
}

//GetName returns the name/filepath of the text file document
func (doc FileDocument) GetName() string {
    return doc.FilePath
}

//QueryDocument document represents a Document constructed from a string query
type QueryDocument struct {
    Query string
    VectorLength float64
    SimilarityScore float64
}

//HashMapVector loops through every term of a query and add terms that
//are not stop words to a vector and returns the vector.
func (doc QueryDocument) HashMapVector() *Vector {
    stopWords := LoadStopWords()
    vector := Vector{make(map[string]float64)}
    terms := strings.Fields(doc.Query) 
    //loop through the terms of the query
    for _, term := range terms {
        strings.ToLower(term)
        //if the term is not a stop word, add it to the vector
        if !stopWords[term] {
            vector.Add(term)
        }
    }
    return &vector
}

/*
//GetVectorLength returns the current value of the vector lenght of a document
func (doc QueryDocument) GetVectorLength() float64 {
    return doc.VectorLength
}

//SetVectorLength sets the value of the vector length of this document to val
func (doc *QueryDocument) SetVectorLength(val float64) {
    doc.VectorLength = val
}

//GetSimilarityScore returns the current value of the similarity score  
func (doc QueryDocument) GetSimilarityScore() float64 {
    return doc.SimilarityScore
}

//SetSimilarityScore sets the value of the similarity score for this document to val
func (doc *QueryDocument) SetSimilarityScore(val float64) {
    doc.SimilarityScore = val
} */


//const stopwords = `a an and are as at be by for from has he in is it its of on that the to was were will with`


//LoadStopWords creates a hash map of stop words.
func LoadStopWords() map[string]bool {
    path, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    baseDir := filepath.Base(path)
    var file *os.File
    if baseDir == "search" {
        file, err = os.Open("stopwords.txt")
    } else {
        file, err = os.Open("../stopwords.txt")
    }
    if err != nil {
        panic(err)
    } 
    scanner := bufio.NewScanner(file)
    stopWords := make(map[string]bool)
    for scanner.Scan() {
        word := scanner.Text()
        fmt.Println(word)
        stopWords[word] = true
    }
    file.Close() 
    /*words := strings.Fields(stopwords)
    stopWords := make(map[string]bool)
    for _, word := range words {
        stopWords[word] = true
    } */
    return stopWords
}

