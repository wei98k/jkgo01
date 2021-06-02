package main

import (
    "fmt"
    "net/http"
    "database/sql"
    "strconv"

    _ "github.com/go-sql-driver/mysql"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    // fmt.Fprintf(w, "hello world")
    query := r.URL.Query()

    ids := query.Get("id")

    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/jkgo01")

    if err != nil {
        panic(err)
    }

    id, err:=strconv.Atoi(ids)

    if err != nil {
        panic(err)
    }
    
    username, err := GetOne(*db, id) 
    
    if err != nil {
        fmt.Fprintf(w, "data is null")
        return
    }
    
    fmt.Fprintf(w, username)
    
}

func main() {

    // init http server
    http.HandleFunc("/test", indexHandler) 
    http.ListenAndServe(":8001", nil) 
}
