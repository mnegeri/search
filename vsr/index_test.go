package vsr

import (
    "testing"
    "path/filepath"
)

func TestProcessDocuments(t *testing.T) {
    dir, _ := filepath.Abs("../test_docs")
    index := InvertedIndex{dir, make([]Document, 0), make(map[string]TermData)}
    index.ProcessDocuments();
    if len(index.IndexedDocs) != 3 {
        t.Errorf("Not all documents were processed")
    }
}

func TestTermHashMapConstruction(t *testing.T) {
    dir, _ := filepath.Abs("../test_docs")
    index := InvertedIndex{dir, make([]Document, 0), make(map[string]TermData)}
    index.ProcessDocuments();
    uniqueTerms := 6
    if len(index.TermHashMap) != uniqueTerms {
        t.Errorf("Ineverted index/TermHashMap incorrectly built")
    }
}

func TestTermHashMapConstructionTwo(t *testing.T) {
    dir, _ := filepath.Abs("../test_docs")
    index := InvertedIndex{dir, make([]Document, 0), make(map[string]TermData)}
    index.ProcessDocuments();
    term := "university"
    expectedTermInstances := 3
    if len(index.TermHashMap[term].InstanceList) != expectedTermInstances {
        t.Errorf("Ineverted index/TermHashMap incorrectly built")
    }
}

func TestRetrieveDocs(t *testing.T) {
    dir, _ := filepath.Abs("../test_docs")
    index := InvertedIndex{dir, make([]Document, 0), make(map[string]TermData)}
    index.ProcessDocuments();
    queryDoc := QueryDocument{Query: "computer science"}
    results := index.RetrieveDocs(queryDoc)
    if len(results) != 3 {
        t.Errorf("RetrieveDocs method not working")
    }
}

func TestInvertedIndex(t *testing.T) {
    dir, _ := filepath.Abs("../test_docs")
    index := InvertedIndex{dir, make([]Document, 0), make(map[string]TermData)}
    index.ProcessDocuments();
    queryDoc := QueryDocument{Query: "computer science"}
    results := index.RetrieveDocs(queryDoc)
    topResult := results[0]
    expectedResult, _ := filepath.Abs("../test_docs/two.txt")
    if topResult.GetName() != expectedResult {
        t.Errorf("Ranking algorithm failed")
    }
}

func TestInvertedIndexTwo(t *testing.T) {
    dir, _ := filepath.Abs("../test_docs")
    index := InvertedIndex{dir, make([]Document, 0), make(map[string]TermData)}
    index.ProcessDocuments();
    queryDoc := QueryDocument{Query: "COMPUTER"}
    results := index.RetrieveDocs(queryDoc)
    topResult := results[0]
    expectedResult, _ := filepath.Abs("../test_docs/two.txt")
    if topResult.GetName() != expectedResult {
        t.Errorf("Ranking algorithm failed")
    }
}



