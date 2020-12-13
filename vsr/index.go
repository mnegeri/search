//Package vsr provides abstractions for developing an information retrieval system using a vector space model.
package vsr 

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "math"
)

//InvertedIndex 
type InvertedIndex struct {
    //Html bool
    Dir string
    IndexedDocs []Document
    TermHashMap map[string]TermData
}

func (index *InvertedIndex) ProcessDocuments() {
    files, err := ioutil.ReadDir(index.Dir)
    if (err != nil) {
        panic(err) 
    }
    for _, file := range files {
        fmt.Println(file.Name())
        doc := Document{filepath.Join(index.Dir, file.Name()), 0,  make(map[string]bool)}
        //vector := doc.HashMapVector()
        index.indexDoc(doc)
    } 
}

func (index *InvertedIndex) indexDoc(doc Document) {
    index.IndexedDocs = append(index.IndexedDocs, doc)
    vector := doc.HashMapVector()
    for term, count := range vector.HashMap {
        var data TermData
        data, present := index.TermHashMap[term]
        if !present {
            data = TermData{InstanceList: make([]TermInstance, 0)}
            index.TermHashMap[term] = data
        } 
        data.InstanceList = append(data.InstanceList, TermInstance{doc, count})
    }
}

//computeIDFandVectorLength computes the tf-idf weigth for every terms, and
//also computes the vector length of every document that was idexed.
func (index *InvertedIndex) computeIDFandVectorLength() {
    numTotalDocs := len(index.IndexedDocs)
    for term, data := range index.TermHashMap {
        //total number of documents containing term
        numTermReferences := len(data.InstanceList)
        //compute idf
        idf := math.Log(float64(numTotalDocs) / float64(numTermReferences))
        if idf != 0 {
            data.Idf = idf
            //update VectorLength of the document in which this term occurs in by (idf * count)^2
            for _, instance := range data.InstanceList {
                instance.Doc.VectorLength += math.Pow(idf * float64(instance.Count), 2)
            }
        } else {
            //if idf is 0, remove term from the inverted index
            delete(index.TermHashMap, term)
        }
    }
    //finish the computation of document vector length by taking the 
    //square root of the value in VectorLength
    for i, doc := range index.IndexedDocs { 
        index.IndexedDocs[i].VectorLength = math.Sqrt(doc.VectorLength)
    }
}
