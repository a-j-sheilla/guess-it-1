package main

import (
	"guess-it-1/student/calculations"
	"strconv"
	"bufio"
	"fmt"
	"os"
	"strings"



)

func main() {
	reader := bufio.NewReader(os.Stdin)

	inputSlice := []string{}

	for {

		// Read input until a newline character
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}
		
		input = strings.TrimSpace(input)
		inputSlice = append(inputSlice, input)
		floatInput := make([]float64, len(inputSlice)-1)
		for _, strnum := range inputSlice{
			num, err := strconv.ParseFloat(strnum, 64)
			if err != nil {
				fmt.Println("error in conversion")
				return
			}
			floatInput= append(floatInput, num)
		}

		std := calculations.StdDev(floatInput)
		Average := calculations.Average(floatInput)
		if len(inputSlice) <= 3 {
			std = 100
		}
		lowerLimit := int(Average) - int(std)*3
		upperLimit := int(Average) + int(std)*3

		// Check for the exit command

		// Print the input back to the user
		fmt.Printf("%d %d\n", lowerLimit, upperLimit)
	}
}
