package vsr

import (
    "testing"
)

func TestLoadStopWords(t *testing.T) {
    doc := Document{"stop_words.txt", "", 0 , 0}
    doc.LoadStopWords()
    _, ok := doc.StopWords["and"]
    if !ok {
        t.Errorf("stop word \"and\" is not present")
    }
}

