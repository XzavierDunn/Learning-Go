package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Calls readLines func to get input into useable format
	lines, err := readLines("day2Input.txt")
	if err != nil {
		fmt.Println(err)
	}

	// Creates a flat string from the array
	justString := strings.Join(lines, " ")
	// Splits flat string up by commas creating another array but less bad
	new := strings.Split(justString, ",")

	// Creates counter var to keep track of postition
	counter := 0
	for index, value := range new {
		_, _ = index, value

		noun := "95"
		verb := "7"
		new[1] = noun
		new[2] = verb

		if counter < len(new) {
			if new[counter] == "1" {
				new = add(new, counter)
			} else if new[counter] == "2" {
				new = mul(new, counter)
			}

			if new[0] == "19690720" {
				fmt.Println(100*stringToInt(noun) + stringToInt(verb))
				os.Exit(3)
			}

			counter = counter + 4

		}
	}

	fmt.Println(new[0])

}

// If "1" call this to add and replace
func add(new []string, counter int) []string {
	// Grabs value at the index
	first := stringToInt(new[counter+1])
	// Grabs the value using the above value
	f := stringToInt(new[first])

	// Grabs value at the index
	second := stringToInt(new[counter+2])
	// Grabs the value using the above value
	s := stringToInt(new[second])

	// Adds up the first and second values and converts back to string
	finalNum := strconv.Itoa(f + s)

	// Gets location which the finalNum should be placed
	location := stringToInt(new[counter+3])

	// Check so you don't get index errors... go doesn't believe in error catching
	if location < len(new) {

		new[location] = finalNum
	}

	return new
}

// If "2" call this to multiply and replace
func mul(new []string, counter int) []string {
	// Grabs value at the index
	first := stringToInt(new[counter+1])
	// Grabs the value using the above value
	f := stringToInt(new[first])

	// Grabs value at the index
	second := stringToInt(new[counter+2])
	// Grabs the value using the above value
	s := stringToInt(new[second])

	// Adds up the first and second values and converts back to string
	finalNum := strconv.Itoa(f * s)

	// Check so you don't get index errors... go doesn't believe in error catching
	location := stringToInt(new[counter+3])

	if location < len(new) {

		new[location] = finalNum
	}

	return new
}

// Used to open and get lines into a readable format
func readLines(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()

}

func stringToInt(x string) (i int) {
	i, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println(err)
	}
	return i
}
