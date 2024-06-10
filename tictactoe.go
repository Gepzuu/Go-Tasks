package main

import "fmt"

var board [3][3]int

func main() {
    InitializeGame()
    GameLoop()
}

func InitializeGame() {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            board[i][j] = 0
        }
    }
}

func GameLoop() {
    for {
        PrintGame()
        MakeMove()
    }
}

func PrintGame() {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            c := ConvertToChar(board[i][j])
            fmt.Printf("%s|", c)
        }
        fmt.Println()
    }
}

func ConvertToChar(num int) string {
    if num == 0 {
        return " "
    } else if num == 1 {
        return "O"
    } else {
        return "X"
    }
}

func MakeMove() {
    var x, y int
    fmt.Print("Enter x: ")
    fmt.Scanln(&x)
    fmt.Print("Enter o: ")
    fmt.Scanln(&y)
    var val int
    fmt.Print("Enter value (1 for O, -1 for X): ")
    fmt.Scanln(&val)
    if PlaceStone(x, y, val) {
        fmt.Println("Move successful!")
    } else {
        fmt.Println("Invalid position to place stone.")
    }
}

func PlaceStone(x, y, val int) bool {
    if x < 0 || x > 2 || y < 0 || y > 2 {
        return false
    }
    if board[x][y] != 0 {
        return false
    }
    board[x][y] = val
    return true
}