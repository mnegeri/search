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

func TestHashMapVector(t *testing.T) {
    doc := Document{"../test_docs/one.txt", "", 0, 0}
    hashMap := doc.HashMapVector()
    count, ok := hashMap.HashMap["university"]
    fmt.Println(count)
    if !ok || (count != 1) {
        t.Errorf("Hash Map Vector not built properly")
    }
}
