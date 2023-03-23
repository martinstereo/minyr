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
func countLines(filename string) int {
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

// Function that adds name of editor of file (me, Martin)
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
	totalTemp := 0.0
	for scanner.Scan() { // scan each line for content

		line := scanner.Text()
		// check if line contains the data
		if strings.Contains(line, "Kjevik;") {
			//take out temp data out of line
			dataArray := strings.Split(line, ";")
			temp := dataArray[3]
			tempFloat, err := strconv.ParseFloat(temp, 64)
			if err != nil {
				log.Fatal(err)
			}
			totalTemp = totalTemp + tempFloat
			countedLines++
		}

	}
	average := totalTemp / float64(countedLines)

	result := fmt.Sprintf("%.2f", average)

	return result
}
