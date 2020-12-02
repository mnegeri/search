package vsr

import (
    "testing"
)

func TestLoadStopWords(t *testing.T) {
    doc := Document{"stop_words.txt"}
    doc.LoadStopWords()
    word, ok := doc.stopWords["and"]
    if !ok {
        t.Errorf("stop word \"and\" is not present")
    }
}
