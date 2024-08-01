# Password Hashing Program

This Go program demonstrates how to hash a password and verify it using the `bcrypt` package from `golang.org/x/crypto/bcrypt`.

## Features

- Hash a plain text password
- Verify a hashed password against a plain text password

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/Gepzuu/Go-Tasks.git
    ```
2. Navigate to the project directory:
    ```sh
    cd your-repo
    ```
3. Install the required package:
    ```sh
    go get golang.org/x/crypto/bcrypt
    ```

## Usage

1. Run the application:
    ```sh
    go run sct.go
    ```
2. Enter a password when prompted.

## Example

```sh
Enter a password: mypassword
Password: mypassword
Hash:     $2a$10$N9qo8uLOickgx2ZMRZo5i.eW8K9h3G8J8K9h3G8J8K9h3G8J8K9h3G8J8K9h3G8
Password match: true
