package main

import (
	"fmt"
)

func main() {

	for {

		fmt.Println("------------------------------------")
		fmt.Println("Welcome to your calculator")
		fmt.Println("------------------------------------")

		var num1, num2 int
		fmt.Print("Enter the first number: ")
		fmt.Scanln(&num1)

		fmt.Print("Enter the second number: ")
		fmt.Scanln(&num2)

		fmt.Println("Choose an operation:")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Division")
		fmt.Println("4. Multiplication")

		var op int
		fmt.Scanln(&op)

		switch op {
		case 1:
			fmt.Println("Result:", num1+num2)
		case 2:
			fmt.Println("Result:", num1-num2)
		case 3:
			if num2 != 0 {
				fmt.Println("Result:", num1/num2)
			} else {
				fmt.Println("Error: Division by zero")
			}
		case 4:
			fmt.Println("Result:", num1*num2)
		default:
			fmt.Println("Invalid operation")
		}
		fmt.Println("Do you want to continue? (y / n)")

		var choice string
		fmt.Scanln(&choice)
		if choice != "y" {
			break

		}
	}
}
