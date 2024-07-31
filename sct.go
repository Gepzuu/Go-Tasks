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