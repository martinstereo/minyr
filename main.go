package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func main() {

	//	buffer := make([]byte, 100)

	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv") // For read access.

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// counts the amount of lines in the file
	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			counter++
		}
	}
	fmt.Println("amount of lines:", counter)


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
}
