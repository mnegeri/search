package main


import ( 
    "fmt"
    //"os"
    "github.com/mnegeri/search/vsr"
)

func main() {
    fmt.Println("test")
    doc := vsr.Document{"hello", 0, make(map[string]bool)}
    doc.LoadStopWords()
    fmt.Println(doc.StopWords["and"])
    /*args := os.Args[1:]
    dir := args[len(args) - 1]
    var stem bool
    var html bool
    for i := 0; i < len(args) - 1; i++ {
        if args[i] == "-stem" {
            stem := true
        } 
        if args[i] == "-html" {
            html := true
        } 
    } */
 //   test := document.Document{"Milki"}
  //  test.Test()
} 
