package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/naausicaa/minyr/yr"
)

// variables for filenames
var (
	inputFile  string = "kjevik-temp-celsius-20220318-20230318.csv"
	outputFile string = "kjevik-temp-fahr-20220318-20230318.csv"
)

func main() {
	fmt.Println("Venligst velg convert, average eller exit:")
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)

			// Convert file
		} else if input == "convert" { // funksjon som gjør åpner fil, leser linjer, gjør endringer og lagrer nye linjer i en ny fil
			// check if file already exists
			if _, err := os.Stat(outputFile); err == nil {
				fmt.Println("converted file already exists. create a new file: y / n")
				// new scanner for checking if user wants to create a new file or not
				var input string
				scanner := bufio.NewScanner(os.Stdin)

				for scanner.Scan() {
					input = scanner.Text()
					// if no - exit
					if input == "n" || input == "no" {
						fmt.Println("exit")
						os.Exit(0)
						// if yes - overwrite file
					} else if input == "y" || input == "yes" {
						//remove existing file before converting again
						err := os.Remove(outputFile)
						if err != nil {
							log.Fatal(err)
						}
						yr.ConvertFile(inputFile)
						fmt.Printf("Lagde fil. Skrev %v linjer.\n", yr.CountLines(outputFile))
						os.Exit(0)
					}
				}
			} else {
				yr.ConvertFile(inputFile)
				fmt.Printf("Lagde fil. Skrev %v linjer.\n", yr.CountLines(outputFile))
				os.Exit(0)
			}
			// averages the temp of file
		} else if input == "average" {
			yr.AverageTempOfFile(inputFile)
		} else {
			fmt.Println("Ugyldig kommando...")
		}
	}

}
