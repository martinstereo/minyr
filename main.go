package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/naausicaa/funtemps/conv" // Pakke som konverterer celsius, fahr og kelvin
)

type yr struct {
	navn    string
	stasjon string
	tid     string
	celsius float64
}

func main() {

	// Open file
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //closes file

	reader := csv.NewReader(bufio.NewReader(file))
	//discard header
	//_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	reader.Comma = ';' // change Comma to the delimiter in the file
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// turn to int
		celsiusFloat, err := strconv.ParseFloat(row[3], 64) // ROW 3 is for celsius / temp
		if err != nil {
			fmt.Println(err)
		}
		conv.CelsiusToFahrenheit(celsiusFloat)

		fmt.Println(row[3])

		// Create new fahrenheit file
		/*
			os.WriteFile("kjevik-temp-fahrenheit-20220318-20230318.csv", []byte(row[3]), 0666)
				if err != nil {
					log.Fatal(err)
				}
			}
		*/
		/*
			// Create a buffered reader for the file
			reader := bufio.NewReader(file)

			// Read text file line by line
			for {
				line, err := reader.ReadString('\n')
				strings.Split(line, ";")
				if err != nil {
					break // reached end of file
				}
				fmt.Println(line)
			}
		*/

		/*
			var input string
			scanner := bufio.NewScanner(os.Stdin)

			for scanner.Scan() {
			    input = scanner.Text()
			    if input == "q" || input == "exit" {
			        fmt.Println("exit")
			        os.Exit(0)
			    } else if input == "convert" {
			        fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Fahrenheit.")
			        // funksjon som gjor aapner fil, leser linjer, gjor endringer og lagrer nye linjer i en ny fil

			    // flere else-if setninger
			    } else {
			        fmt.Println("Venligst velg convert, average eller exit:")

			    }

			}
		*/
		// counts the amount of lines in the file

		/*
			scanner := bufio.NewScanner(file)
			counter := 0
			for scanner.Scan() {
				line := scanner.Text()
				if line != "" {
					counter++
				}
			}
			fmt.Println("total lines:", counter)
		*/

		//fahr := conv.CelsiusToFahrenheit(100)
		//fmt.Println("100C til F Er", fahr)
	}
}
