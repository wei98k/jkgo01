package main

import (
    "fmt"
)

func Positive(n int) *bool {
    if n == 0 {
        return nil
    }
    
    r := n > -1
    return &r
}

func Check(n int) {
    pos := Positive(n)
    if pos == nil {
        fmt.Println(n, "is neither")
        return
    }
    if *pos {
        fmt.Println(n, "is positive")
    } else {
        fmt.Println(n, "is negative")
    }
}

func main() {
   Check(1) 
   Check(0) 
   Check(-1) 
}
