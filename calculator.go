//Simple Calculator

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <number> <operator> <number>")
		return
	}

	//Validate input arguments
	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid Input number 1:", err)
		return
	}

	operator := os.Args[2]

	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Invalid Input number 2:", err)
		return
	}

	// Perform arithmetic based on the operator
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: division by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operator:", operator)
		return
	}

	// Print the result
	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}

/*

# Output

$ go run calculator.go 50 + 38
50.00 + 38.00 = 88.00

*/