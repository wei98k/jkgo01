package main

import (
    "fmt"
    "net/http"
    "log"
    "os"
    "os/signal"
)

func main() {
    
    c := make(chan os.Signal)
    signal.Notify(c)

    fmt.Println("start .....")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Are you ok? ")
    })
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
    // 如何注册linux signal 信号和处理?
    
    s := <-c

    fmt.Println("End...", s)
}
