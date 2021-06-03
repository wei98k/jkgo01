package main

import (
    "fmt"
)

func Positive(n int) bool {
    if n == 0 {
        panic("undefined")
    }
    return n > -1
}

func Check(n int) {
    defer func() {
        if recover() != nil {
            fmt.Println("is neither")
        }
    }()
    if Positive(n) {
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
