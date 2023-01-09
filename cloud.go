package main

import (
    "log"
    "database/sql"
)

// Cloud represents a connection to the cloud database
type Cloud struct {
    db *sql.DB
}

// Connect establishes a connection to the cloud database
func (c *Cloud) Connect() error {
    // Connect to the database
    db, err := sql.Open("mysql", "user:password@/dbname")
    if err != nil {
        return err
    }
    c.db = db

    // Ping the database to ensure the connection is working
    err = c.db.Ping()
    if err != nil {
        return err
    }

    return nil
}

// Close closes the connection to the cloud database
func (c *Cloud) Close() {
    err := c.db.Close()
    if err != nil {
        log.Println(err)
    }
}
