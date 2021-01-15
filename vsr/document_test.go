package vsr

import (
    "testing"
    "fmt"
)

func TestLoadStopWords(t *testing.T) {
    stopWords := LoadStopWords()
    _, ok := stopWords["and"]
    if !ok {
        t.Errorf("stop word \"and\" is not present")
    }
}

func TestStopWordsSize(t *testing.T) {
    stopWords := LoadStopWords()
    numStopWords := 127
    if len(stopWords) != numStopWords {
        t.Errorf("stop words hashmap not loaded properly")
    }
}


func TestHashMapVector(t *testing.T) {
    doc := FileDocument{"../test_docs/one.txt", 0, 0}
    hashMap := doc.HashMapVector()
    count, ok := hashMap.HashMap["university"]
    fmt.Println(count)
    if !ok || (count != 1) {
        t.Errorf("Hash Map Vector not built properly")
    }
}

func TestGetName(t *testing.T) {
    path := "/home/pretend/path"
    doc := FileDocument{FilePath: path}
    if doc.GetName() != path {
        t.Errorf("GetName method not working")
    }
}

func TestQueryHashMapVector(t *testing.T) {
    doc := QueryDocument{Query: "This is the query for search"}
    vector := doc.HashMapVector()
    if len(vector.HashMap) != 2 {
        t.Errorf("Hash Map Vector properly incorrectly formed from the query")
    }
}

func TestQueryHashMapVectorTwo(t *testing.T) {
    doc := QueryDocument{Query: "query Query query Query query Query"}
    vector := doc.HashMapVector()
    termFreq := float64(6)
    if vector.HashMap["query"] != termFreq || len(vector.HashMap) != 1 {
        t.Errorf("Hash Map Vector properly formed from a QueryDocument")
    }
}




