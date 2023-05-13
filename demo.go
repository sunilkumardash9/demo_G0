package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Addition Program")
	fmt.Println("----------------")

	// Read the first number
	fmt.Print("Enter the first number: ")
	num1Str, _ := reader.ReadString('\n')
	num1Str = strings.TrimSpace(num1Str)
	num1, _ := strconv.Atoi(num1Str)

	// Read the second number
	fmt.Print("Enter the second number: ")
	num2Str, _ := reader.ReadString('\n')
	num2Str = strings.TrimSpace(num2Str)
	num2, _ := strconv.Atoi(num2Str)

	result := Add(num1, num2)

	fmt.Printf("Sum: %d\n", result)
}

// Add function takes two integers and returns their sum.
func Add(a, b int) int {
	return a + b
}
