package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	help := flag.Bool("Help", false, "Show more Information\n The wordc gives output of the number of words in the documednt")

	flag.Parse()
	if *help {
		flag.Usage()
		return
	}
	// Check if a filename is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize counters
	lineCount := 0
	wordCount := 0
	charCount := 0

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text()

		// Count the number of lines
		lineCount++

		// Count the number of words
		words := strings.Fields(line)
		wordCount += len(words)

		// Count the number of characters
		charCount += len(line) + 1 // Add 1 for newline character
	}

	// Check if there was any error during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// Print the results
	fmt.Printf("Lines: %d\n", lineCount)
	fmt.Printf("Words: %d\n", wordCount)
	fmt.Printf("Characters: %d\n", charCount)
}
