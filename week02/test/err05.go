package main

import (
    "errors"
    "fmt"
)

func Positive(n int) (bool, error) {
    if n == 0 {
        return false, errors.New("undefined")
    }
    return n > -1, nil
}

func Check(n int) {
    pos , err := Positive(n)
    if err != nil {
        fmt.Println(n, err)
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
