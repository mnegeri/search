package vsr



type TermData struct {
    Idf uint64
    InstanceList []TermInstance
}


type TermInstance struct {
    Count int
}
