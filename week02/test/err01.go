package main

import (
    "errors"
    "fmt"
)

// Create a named type for our new error type 
type errorString string

// Implement the error interface.
func (e errorString) Error() string {
    return string(e)
}

// New creates interface values of type error
func New(text string) error {
    return errorString(text)
}

var ErrNamedType = New("EOF")
var ErrStructType = errors.New("EOF")


func main() {
    if ErrNamedType == New("EOF") {
        fmt.Println("Named Type Error")
    }

    if ErrStructType == errors.New("EOF") {
        fmt.Println("Struct Type Error")
    }
        
}



