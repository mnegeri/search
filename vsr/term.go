package vsr


//TermData represents information on all instances  where a certain 
//term occurs in the corpus and the computed idf value of that term.
type TermData struct {
    InstanceList []TermInstance
    Idf float64 
}

//TermInstance represent a single instance of a term. Stores 
//the document is occurs in and the number of times it occurs.
type TermInstance struct {
    Doc Document
    Count float64 
}
