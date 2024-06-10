package main

import (
	"crypto/rand"
	"math/big"
	"fmt"
)

// charset contains the characters that can be used in the generated password
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	passwordLength := 16 // Set the length of the generated password
	password := GeneratePassword(passwordLength) // Generate the password
	fmt.Println("Generated password:", password) // Print the generated password
}

// GeneratePassword generates a random password of the specified length using characters from the charset
func GeneratePassword(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; {
		// Generate a random number within the range of the charset length
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			fmt.Println(err)
			return ""
		}
		// Use the random number as an index to select a character from the charset
		b[i] = charset[n.Int64()]
		i++
	}
	return string(b) // Convert the byte slice to a string and return the generated password
}

/*

# Output

Generated password: dhrCW6pNKpEvVCR

*/

