package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Not Found:", err)
		return
	}
	defer file.Close() // Ensure the file is closed

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Slice to store numbers
	var numbers []float64

	// Read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into fields (words/numbers)
		for _, field := range strings.Fields(line) {
			// Convert the field to a float
			num, err := strconv.ParseFloat(field, 64)
			if err != nil {
				fmt.Println("Error parsing number:", field, err)
				continue
			}
			numbers = append(numbers, num)
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Calculate and print the average if numbers are available
	if len(numbers) > 0 {
		avg := Median(numbers)
		fmt.Printf("Median: %.2f\n", avg)
	} else {
		fmt.Println("No numbers found to calculate the median.")
	}
}




func Variance(input []float64)float64{}