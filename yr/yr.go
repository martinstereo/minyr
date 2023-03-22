package yr

import (
	"bufio"
	"log"
	"os"
)

// function that counts the amout of lines in a file
func countLines(inputFile string) int {
	// open file
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()                //closes file
	scanner := bufio.NewScanner(file) // create scanner from bufio package
	countedLines := 0                 // intitale variable with amount of lines
	for scanner.Scan() {              // scan each line for content
		line := scanner.Text()
		if line != "" {
			countedLines++
		}
	}
	return countedLines
}

func ConvertCelsiusToFahr(input string) string {

	return ""
}
