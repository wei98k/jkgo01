package main

import (
	"database/sql"
    "github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
)


func GetOne(db sql.DB,id int) (string, error) {

    var username string

    err := db.QueryRow("select username from users where id = ?", id).Scan(&username)
    if err != nil {
        if err == sql.ErrNoRows {
            //fmt.Println("sql.ErrNoRows")
            return "", errors.Wrap(err, "data is nil")
        }
        return "", err
    } 
    return username, nil 
}
