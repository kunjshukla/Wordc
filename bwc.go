package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var filename string
	var char, lin, word bool

	flag.BoolVar(&char, "c", false, "Count characters only")
	flag.BoolVar(&lin, "l", false, "Count lines only")
	flag.BoolVar(&word, "w", false, "Count words only")

	flag.StringVar(&filename, "f", "", "Target File")

	flag.Parse()

	if !checkFileExists(filename) {
		fmt.Printf("Could not find/read %s\n", filename)
		os.Exit(1)
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %s\n", err)

	}
	defer file.Close()

	if !(char || lin || word) {
		fmt.Println(getCharacters(file), getWords(file), getLines(file), file.Name())
		os.Exit(0)
	}

	if char {
		fmt.Print(getCharacters(file), " ")
	}

	if lin {
		fmt.Print(getLines(file), " ")
	}

	if word {
		fmt.Print(getWords(file), " ")
	}

	fmt.Println(file.Name())
}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func getCharacters(file *os.File) int64 {
	file.Seek(0, 0)
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Error getting file stats: %s\n", err)
		return 0
	}
	return fileInfo.Size()
}

func getLines(file *os.File) int {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	return lineCount
}

func getWords(file *os.File) (wordCount int) {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordCount++
	}

	return
}
