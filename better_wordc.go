package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var filename string
	var help bool

	flag.BoolVar(&help, "h", false, "Show more Information\n The wordc gives output of the number of words in the documednt")
	flag.StringVar(&filename, "f", "", "Target File")
	// Other command line args
	// ...

	flag.Parse()

	if help {
		printHelp()
		os.Exit(0)
	}

	if !checkFileExists(filename) {
		log.Fatalf("Could not find/read %s\n", filename)
	}

	// Read file here or get the scanner object to pass into the getWords/getLines functions

	charCount := getCharacters("/home/hp/Wordc/LICENSE")
	// lineCount := getLines()
	// wordCount := getWords()

	// fmt.Println(wordCount, lineCount, charCount)
	fmt.Println(charCount)

}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}

func printHelp() {

}

func getCharacters(filename string) uint64 {
	// use something like os.File.Stat form stdlib
	file, _ := os.Stat(filename)
	return uint64(file.Size())
}

func getLines(buffer *bufio.Reader) uint64 {

	return 0
}

func getWords(filename *bufio.Reader) uint64 {
	return 0
}
