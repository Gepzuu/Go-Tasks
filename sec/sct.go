package main

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

// Function to hash a password
func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// Function to compare a hashed password with a plain password
func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}


func main() {
    password := "mySecurePassword"
    hash, err := hashPassword(password)
    if err != nil {
        fmt.Println("Error hashing password:", err)
        return
	}

    fmt.Println("Password:", password)
    fmt.Println("Hash:    ", hash)

    match := checkPasswordHash(password, hash)
    fmt.Println("Password match:", match)
}