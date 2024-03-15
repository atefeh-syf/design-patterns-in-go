package main

import (
 "database/sql"
 "errors"
 "fmt"
 "time"
)

type Database interface {
 Connect() (*sql.DB, error)
}

type RealDatabase struct{}

func (r *RealDatabase) Connect() (*sql.DB, error) {
 db, err := sql.Open("sqlite3", "test.db")
 if err != nil {
  return nil, err
 }
 time.Sleep(2 * time.Second)
 return db, nil
}

type ProxyDatabase struct {
 realDatabase *RealDatabase
}

func (p *ProxyDatabase) Connect() (*sql.DB, error) {
 if p.realDatabase == nil {
  p.realDatabase = &RealDatabase{}
 }

 if !isUserAuthorized() {
  return nil, errors.New("user is not authorized to access the database")
 }

 fmt.Println("User accessed the database at", time.Now().Format(time.RFC3339))
 return p.realDatabase.Connect()
}

func isUserAuthorized() bool {
 return true
}

func main() {
 var db Database = &ProxyDatabase{}
 connection, err := db.Connect()
 if err != nil {
  fmt.Println("Failed to connect to the database:", err)
  return
 }
 defer connection.Close()
 fmt.Println("Successfully connected to the database.")
}