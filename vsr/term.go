package vsr


type TermData struct {
    InstanceList []TermInstance
    Idf uint64
}


type TermInstance struct {
    Doc Document
    Count int
}
