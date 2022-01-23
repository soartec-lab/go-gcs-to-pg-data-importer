package main

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

// メイン関数
func main() {
    db, err := sql.Open("postgres", "host=postgresql port=5432 user=postgres sslmode=disable")
    defer db.Close()

    if err != nil {
        fmt.Println(err)
    }

    // INSERT
    res, err := db.Exec("INSERT INTO users (id, name, age) VALUES (1, 'tom', 18)")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("%+v", res)

    // SELECT
    rows, err := db.Query("SELECT * FROM users")

    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("%+v", *rows)
}