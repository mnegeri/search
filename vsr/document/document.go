package document 

import (
   // "io/ioutil"
    //"os"
    "fmt"
)

type Document struct {
    Name string 
}

func (doc Document) Test() {
    fmt.Println(doc.Name)
}
