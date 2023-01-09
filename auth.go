package main

import (
    "fmt"
    "log"
    "database/sql"
    "golang.org/x/crypto/bcrypt"
)

// User represents a user of the web GUI
type User struct {
    ID int
    Username string
    PasswordHash []byte
}

// Connect establishes a connection to the database
func Connect() (*sql.DB, error) {
    // Connect to the database
    db, err := sql.Open("mysql", "user:password@/dbname")
    if err != nil {
        return nil, err
    }

    // Ping the database to ensure the connection is working
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return db, nil
}

// GetUser gets a user from the database by username
func GetUser(db *sql.DB, username string) (*User, error) {
    // Execute a query to get the user with the given username
    rows, err := db.Query("SELECT id, username, password_hash FROM users WHERE username = ?", username)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    // Scan the result into a user
    var user User
    if rows.Next() {
        err = rows.Scan(&user.ID, &user.Username, &user.PasswordHash)
        if err != nil {
            return nil, err
        }
    } else {
        return nil, fmt.Errorf("User not found")
    }

    return &user, nil
}

// CheckPassword checks if a password is correct for a given user
func CheckPassword(user *User, password string) bool {
    // Compare the password hash to the given password
    err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
    if err != nil {
        log.Println(err)
        return false
    }
    return true
}
