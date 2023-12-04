package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

const (
	FILEPATH   = "./resources/text_input.txt"
	MAX_LENGTH = 5
)

func main() {
	// Define variables
	var runningSum int = 0
	var digitMap map[string]int

	digitMap = make(map[string]int)
	digitMap["one"] = '1'
	digitMap["two"] = '2'
	digitMap["three"] = '3'
	digitMap["four"] = '4'
	digitMap["five"] = '5'
	digitMap["six"] = '6'
	digitMap["seven"] = '7'
	digitMap["eight"] = '8'
	digitMap["nine"] = '9'

	// Read file
	file, err := os.Open(FILEPATH)
	if err != nil {
		log.Fatalf(err.Error())
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Iterate over each line of the file and iterate over each byte in the line
	for scanner.Scan() {
		var numSlice []rune
		var currentString string
		p1 := 0
		p2 := p1
		line := scanner.Text()

		for p1 < len(line) {
			if len(currentString) < MAX_LENGTH && p2 < len(line) {
				p1Char := rune(line[p1])
				p2Char := rune(line[p2])

				if p1 == p2 && unicode.IsDigit(p1Char) {
					numSlice = append(numSlice, p1Char)
					p1++
					p2 = p1
					currentString = ""
					continue
				} else if p1 != p2 && unicode.IsDigit(p2Char) {
					p1++
					p2 = p1
					currentString = ""
					continue
				}

				currentString = currentString + string(p2Char)

				if value, ok := digitMap[currentString]; ok {
					byteValue := byte(value)
					numSlice = append(numSlice, rune(byteValue))
					p1 = p2 + 1
					p2 = p1
					currentString = ""
				} else {
					p2++
				}

			} else {
				currentString = ""
				p1++
				p2 = p1
			}
		}

		firstNumber := string(numSlice[0])
		lastNumber := string(numSlice[len(numSlice)-1])
		numberToAdd := firstNumber + lastNumber

		num, err := strconv.ParseInt(numberToAdd, 10, 16)
		if err != nil {
			log.Fatalf("Couldnt convert %s", numberToAdd)
		}
		fmt.Printf("Number is %d\n", num)

		runningSum += int(num)
	}

	// return running sum
	fmt.Printf("Sum is %d\n", runningSum)
}
