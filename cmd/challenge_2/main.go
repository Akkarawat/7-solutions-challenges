package main

import (
	minrl "7-solutions-challenges/internal/min_RL"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter input (Press Ctrl+C to exit):")

	for {
		fmt.Print("> ") // Command-line prompt
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			fmt.Println("Input cannot be empty. Please enter again.")
			continue
		}

		// Process the input using SolveMinRL
		result := minrl.SolveMinRL(input)
		fmt.Println("Result:", result)
	}
}