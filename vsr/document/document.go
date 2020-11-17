package document 

import (
   // "io/ioutil"
    //"os"
    "fmt"
)

type Document struct {
    name string 
}

func (doc Document) Test() {
    fmt.Println(doc.name)
}
