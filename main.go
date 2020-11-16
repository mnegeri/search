package main

import ( 
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]
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
    }
}
