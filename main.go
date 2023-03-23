package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/naausicaa/minyr/yr"
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
		} else if input == "convert" {
			// funksjon som gjør åpner fil, leser linjer, gjør endringer og lagrer nye linjer i en ny fil
			fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Fahrenheit.")
			// Open file
			file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv") // For read access.
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close() //closes file

			//create scanner
			scanner := bufio.NewScanner(file)
			scanner.Scan() // move to next token before loop to skip first line
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Println(yr.ConvertCelsiusToFahr(line))
			}
			if err != nil {
				log.Fatal(err)
			}

			/*
				//create new file
				outputFile, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv") //create output file for fahr
				if err != nil {
					log.Fatal(err)
				}
				//close new file
				defer outputFile.Close() // closes output file
				if err != nil {
					log.Fatal(err)
				}

				//writer := bufio.NewWriter(outputFile)

			*/

			// averages the temp of file
		} else if input == "average" {
			avg := yr.AverageTemp("kjevik-temp-celsius-20220318-20230318.csv")
			fmt.Println("average celsius temperature of period is ", avg)
		}
	}
}
