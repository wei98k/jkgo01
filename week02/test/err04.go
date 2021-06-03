package main

import "fmt"

// Positive return true if the number is positive, false if it is negative.
func Positive(n int) (bool, bool) {
    if n == 0 {
        return false, false
    }
    return n > -1, true
}


func Check(n int) {
    pos, ok := Positive(n)
    if !ok {
        fmt.Println(n, "is neither")
        return 
    }
    if pos {
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

