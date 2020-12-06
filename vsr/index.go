package vsr 

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
)

type InvertedIndex struct {
    //Html bool
    Dir string
    TermHashMap map[string]TermData
}

func (i *InvertedIndex) ProcessDocuments() {
    files, err := ioutil.ReadDir(i.Dir)
    if (err != nil) {
        panic(err) 
    }
    for _, file := range files {
        fmt.Println(file.Name())
        doc := Document{filepath.Join(i.Dir, file.Name()), make(map[string]bool)}
        //vector := doc.HashMapVector()
        i.indexDoc(doc)
    } 
}

func (i *InvertedIndex) indexDoc(doc Document) {
    vector := doc.HashMapVector()
    for term, count := range vector.HashMap {
        var data TermData
        data, present := i.TermHashMap[term]
        if !present {
            data = TermData{InstanceList: make([]TermInstance, 0)}
            i.TermHashMap[term] = data
        } 
        append(data.InstanceList, TermInstance{doc, count})
    }
}
