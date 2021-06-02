package main

import (
    "fmt"

    "github.com/pkg/errors"
)

func login(apikey string) error {
    return errors.New("demoapp")
}

type stackTracer interface {
    StackTrace() errors.StackTrace
}

func main() {
    err, ok := errors.Cause(login("demoapp")).(stackTracer)
    if !ok {
        panic("oops, err does not implement stackTracer")
    }
    st := err.StackTrace()
    fmt.Printf("%+v", st[0:]) 
}
