package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Database connection string
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
    
    // Open a connection to the database
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Check if the connection is successful
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected to the database!")
}