package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

const FILEPATH = "./resources/text_input.txt"

func main() {

	// Create variables
	var runningSum int = 0

	// Load the file
	file, err := os.Open(FILEPATH)
	if err != nil {
		log.Fatalf(err.Error())
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Iterate over each character on each line of the file
	for scanner.Scan() {
		line := scanner.Text()
		var acc []rune
		for _, char := range line {
			if unicode.IsDigit(char) {
				acc = append(acc, char)
			}
		}

		first := string(acc[0])
		last := string(acc[len(acc)-1])

		valueToAdd := first + last
		fmt.Printf("Adding %s and %s\n", first, last)

		// Add to the running sum
		number, err := strconv.ParseInt(valueToAdd, 10, 16)
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Printf("Number is %d\n", number)
		runningSum += int(number)
	}

	// return running sum
	fmt.Printf("Sum is %d\n", runningSum)
}
