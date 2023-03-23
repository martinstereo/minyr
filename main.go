package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/naausicaa/funtemps/conv"
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
		} else if input == "convert" { // funksjon som gjør åpner fil, leser linjer, gjør endringer og lagrer nye linjer i en ny fil
			// check if file already exists
			if _, err := os.Stat("kjevik-temp-fahr-20220318-20230318.csv"); err == nil {
				fmt.Println("File already exists. Overwrite existing file: y / n")

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
						convertFile()
					}
				}
			} else {
				convertFile()
			}
			// averages the temp of file
		} else if input == "average" {
			averageTemp()
		}
	}
}

func convertFile() {
	fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Fahrenheit.")
	// Open file
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //closes file

	//create scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//create output file for fahr
		outputFile, err := os.OpenFile("kjevik-temp-fahr-20220318-20230318.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}
		//close new file
		defer outputFile.Close() // closes output file

		writer := bufio.NewWriter(outputFile)

		if strings.Contains(line, "Navn") {
			writer.WriteString(line + "\n")
		} else if strings.Contains(line, "Kjevik;") {
			writer.WriteString(yr.ConvertCelsiusToFahr(line) + "\n")
		} else if strings.Contains(line, "Data") {
			writer.WriteString(yr.EditEndLine(line) + "\n")
		}
		writer.Flush()
	}
}

func averageTemp() {
	fmt.Println("Venligst velg mellom celsius eller fahrenheit: c / f")
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
		avgCelsius := yr.AverageTemp("kjevik-temp-celsius-20220318-20230318.csv")
		// give user ability to exit
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)

		} else if input == "c" { // celsius
			fmt.Printf("average celsius temperature of period is %v°C\n", avgCelsius)

		} else if input == "f" { // fahrenheit
			//convert to float before converting
			celsiusFloat, err := strconv.ParseFloat(avgCelsius, 64)
			if err != nil {
				log.Fatal(err)
			}
			avgFahr := conv.CelsiusToFahrenheit(celsiusFloat)
			fmt.Printf("Average temperature of period in file is %.2f°F\n", avgFahr)
		}
	}
}
