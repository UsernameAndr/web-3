package main

import "fmt"

func main() {
    var myString string
    fmt.Scanln(&myString) 
    
    var mxrune rune = 0 
    for _, v := range myString {
        if v > mxrune {
            mxrune = v
        }
    }
    
    fmt.Printf("%c", mxrune)
}
