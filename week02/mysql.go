package main

import (
	"database/sql"
    "fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ...

type Users struct {
    id int64
    username string
    password string
}


func main() {
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/jkgo01")
    if err != nil {
        panic(err)
    } 

    defer db.Close()

    var username string
    
    err = db.QueryRow("select username from users where id = ?", 2).Scan(&username)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("sql.ErrNoRows") 
        } else {
           fmt.Println("sql-error") 
        }
    }
    fmt.Println(username)
}
