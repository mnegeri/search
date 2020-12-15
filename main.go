package main


import ( 
    "fmt"
    "os"
    "github.com/mnegeri/search/vsr"
    "bufio"
)

func main() {
    /* fmt.Println("test")
    doc := vsr.Document{"hello", 0, make(map[string]bool)}
    doc.LoadStopWords()
    fmt.Println(doc.StopWords["and"]) */
    args := os.Args[1:]
    dir := args[len(args) - 1]
    cont := true
    reader := bufio.NewReader(os.Stdin)
    index := vsr.InvertedIndex{dir, make([]vsr.Document, 0), make(map[string]vsr.TermData)}
    index.ProcessDocuments()
    for cont {
        fmt.Println("Enter search query: ")
        query, _ := reader.ReadString('\n')
        results := index.RetrieveDocs(&vsr.Document{"", query, 0, 0})
        vsr.PrintResults(results)
        fmt.Println("Press y to continue: ")
        resp, _ := reader.ReadString('\n')
        if resp == "y\n" {
            cont = false
        }
    } 
} 
