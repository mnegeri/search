package vsr

import (
    "testing"
    "path/filepath"
)

func TestInvertedIndex(t *testing.T) {
    dir, _ := filepath.Abs("../test_docs")
    index := InvertedIndex{dir, make([]Document, 0), make(map[string]TermData)}
    index.ProcessDocuments();
    if len(index.IndexedDocs) != 3 {
        t.Errorf("Not all documents were processed")
    }
}

