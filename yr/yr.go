package yr

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/naausicaa/funtemps/conv"
	"github.com/naausicaa/funtemps/format"
)

// function that counts the amout of lines in a file
func CountLines(filename string) int {
	file, err := os.Open(filename) // open file
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

// function that takes lines from yr file and converts temp from celsius to fahrenheit
func ConvertCelsiusToFahr(inputLine string) string {
	// Navn;Stasjon;Tid(norsk normaltid);Lufttemperatur
	var yrData struct {
		navn    string
		stasjon string
		tid     string
		temp    string
	}
	dataArray := strings.Split(inputLine, ";") // split string by semicolon
	// store data values in variables
	yrData.navn = dataArray[0]
	yrData.stasjon = dataArray[1]
	yrData.tid = dataArray[2]
	yrData.temp = dataArray[3]

	// convert to float64 in order to convert to fahr
	celsius, err := strconv.ParseFloat(yrData.temp, 64)
	if err != nil {
		log.Fatal(err)
	}
	//Convert to fahrenheit and format to string from funtemps - conv and format
	yrData.temp = format.FormatOutput(conv.CelsiusToFahrenheit(celsius))
	// restructure the original string
	newLine := []string{
		yrData.navn,
		yrData.stasjon,
		yrData.tid,
		yrData.temp,
	}
	convertedString := strings.Trim(strings.Join(newLine, ";"), "[]{}")
	return convertedString
}

// Function that adds name of editor of file to last line of specific file (me, Martin)
func EditEndLine(lastLine string) string {
	lineArray := strings.Split(lastLine, ";")
	copyright := lineArray[0]
	lineArray[1] = "endringen er gjort av Martin Steiro"
	newEndLine := []string{
		copyright,
		lineArray[1],
	}
	convertedEndLine := strings.Trim(strings.Join(newEndLine, ";"), "[]{}")

	return convertedEndLine
}

// function that calculates average temperature of period of file as a string
func AverageTemp(filename string) string {
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //closes file

	scanner := bufio.NewScanner(file) // create scanner from bufio package
	countedLines := 0                 // intitale variable with amount of lines
	totalTemp := 0.0                  // variable for total temperature added up
	for scanner.Scan() {              // scan each line for content
		line := scanner.Text()
		// check if line contains the data
		if strings.Contains(line, "Kjevik;") {
			//take out temperature data out of line
			dataArray := strings.Split(line, ";")
			temp := dataArray[3]
			tempFloat, err := strconv.ParseFloat(temp, 64) //convert data to float - orignally a string
			if err != nil {
				log.Fatal(err)
			}
			totalTemp = totalTemp + tempFloat // add together
			countedLines++                    // keeping count of lines of data
		}
	}
	average := totalTemp / float64(countedLines) // calculate average by dividing total temp by counted lines of data
	result := fmt.Sprintf("%.2f", average)       // and turn to string with 2 decimal format

	return result
}

func ConvertFile(inputFile string) {
	fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Fahrenheit.")
	// Open file
	file, err := os.Open(inputFile) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //closes file

	//open new file, give read write access
	outputFile, err := os.OpenFile("kjevik-temp-fahr-20220318-20230318.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close() // closes output file
	//create scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//create output file for fahr
		// create writer
		writer := bufio.NewWriter(outputFile)

		// only convert lines with data that need change
		if strings.Contains(line, "Navn") {
			writer.WriteString(line + "\n")
		} else if strings.Contains(line, "Kjevik;") {
			writer.WriteString(ConvertCelsiusToFahr(line) + "\n")
		} else if strings.Contains(line, "Data") {
			writer.WriteString(EditEndLine(line) + "\n") // adds "endring gjort av Martin" to last line
		}
		// flush
		writer.Flush()
	}
}

func AverageTempOfFile(file string) {
	fmt.Println("Venligst velg mellom celsius eller fahrenheit: c / f")
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
		avgCelsius := AverageTemp(file)
		// give user ability to exit
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)

		} else if input == "c" { // celsius
			fmt.Printf("average celsius temperature of period is %v°C\n", avgCelsius)

		} else if input == "f" { // fahrenheit
			//convert to float before converting to fahrenheit
			celsiusFloat, err := strconv.ParseFloat(avgCelsius, 64)
			if err != nil {
				log.Fatal(err)
			}
			avgFahr := conv.CelsiusToFahrenheit(celsiusFloat)
			fmt.Printf("Average temperature of period in file is %.2f°F\n", avgFahr)
		}
	}
}
