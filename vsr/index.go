//Package vsr provides structures and methods for developing an 
//information retrieval system using a vector space model.
package vsr 

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "math"
    "sort"
)

//InvertedIndex represents an inverted index constructed from documents in a directory. 
type InvertedIndex struct {
    //Html bool
    Dir string
    IndexedDocs []Document
    TermHashMap map[string]TermData
}

//ProcessDocuments loops through documents in a directory and builds an inverted index.
func (index *InvertedIndex) ProcessDocuments() {
    files, err := ioutil.ReadDir(index.Dir)
    if (err != nil) {
        panic(err) 
    }
    //loop through each file document in the directory
    for _, file := range files {
        fmt.Println(file.Name())
        doc := &FileDocument{FilePath: filepath.Join(index.Dir, file.Name())}
        //vector := doc.HashMapVector()
        index.indexDoc(doc)
    } 
}

//indexDoc adds a document to the inverted index
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
    index.computeIDFandVectorLength()
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
                instance.Doc.SetSimilarityScore(instance.Doc.GetVectorLength() + math.Pow(idf * float64(instance.Count), 2))
                //instance.Doc.VectorLength += math.Pow(idf * float64(instance.Count), 2)
            }
        } else {
            //if idf is 0, remove term from the inverted index
            delete(index.TermHashMap, term)
        }
    }
    //finish the computation of document vector length by taking the 
    //square root of the value in VectorLength
    for i, doc := range index.IndexedDocs { 
        //index.IndexedDocs[i].VectorLength = math.Sqrt(doc.VectorLength)
        index.IndexedDocs[i].SetVectorLength(math.Sqrt(doc.GetVectorLength()))
    }
}


//RetrieveDocs takes a query and return an array of documents in the 
//order of their relevance.
func (index *InvertedIndex) RetrieveDocs(query QueryDocument) []Document {
    queryVector := query.HashMapVector()
    //hashmap to store the documents that are retrieved and their partially accumulated
    //scores(cosine similarity)
    retrievals := make(map[Document]float64)
    queryLength := 0.0
    for term, count := range queryVector.HashMap {
        queryLength += index.addTermToIndex(term, retrievals, count) 
    }
    queryLength = math.Sqrt(queryLength)
    result := make([]Document, 0)

    for doc, value := range retrievals {
        //doc.SimilarityScore = value / (queryLength * doc.VectorLength) 
        //normalize similarity score
        doc.SetSimilarityScore(value / (queryLength * doc.GetVectorLength()))
        result = append(result, doc) 
    }
    sort.Slice(result, func(i, j int) bool {
        return result[i].GetSimilarityScore() < result[j].GetSimilarityScore()
    })
    return result
}

func (index *InvertedIndex) addTermToIndex(term string, retrievals map[Document]float64, count float64) float64 {
    termData, present := index.TermHashMap[term]
    if present {
        tfIdf := termData.Idf * count
        //loop through each document instance in which the term occurs
        for _, instance := range termData.InstanceList {
            _, present :=  retrievals[instance.Doc]
            if !present {
                retrievals[instance.Doc] = 0
            }
            retrievals[instance.Doc] += tfIdf * termData.Idf * instance.Count
        }
        return tfIdf * tfIdf
    } else {
        return 0.0
    }
}

func PrintResults(results []Document) {
    for _, doc := range results {
        fmt.Println(doc.GetName())
    }
}
