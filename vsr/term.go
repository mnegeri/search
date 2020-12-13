package vsr


type TermData struct {
    InstanceList []TermInstance
    Idf float64 
}


type TermInstance struct {
    Doc Document
    Count int
}
