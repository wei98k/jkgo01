package main

import (
    "fmt"
)

type errorString struct {
    s string
}

func (e errorString) Error() string {
    return e.s
}

func NewError(text string) error {
    return errorString{text}
}

var ErrType = NewError("EOF")

func main() {
    if ErrType == NewError("EOF") {
        fmt.Println("Error:", ErrType)
    }
}
