package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
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
		// Convert the line to a float
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error parsing number:", line, err)
			continue
		}
		numbers = append(numbers, num)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Calculate and print statistics if numbers are available
	if len(numbers) > 0 {
		average := Average(numbers)
		median := Median(numbers)
		variance := Variance(numbers)
		stdDev := math.Sqrt(variance)

		fmt.Println("Average:", int(math.Round(average)))
		fmt.Println("Median:", int(math.Round(median)))
		fmt.Println("Variance:", int(math.Round(variance)))
		fmt.Println("Standard Deviation:", int(math.Round(stdDev)))
	} else {
		fmt.Println("No numbers found to calculate statistics.")
	}
}

// Average calculates the mean of a slice of numbers
func Average(numbers []float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

// Variance calculates the variance of a slice of numbers
func Variance(input []float64) float64 {
	if len(input) == 0 {
		return 0 // Handle empty input case
	}

	// Step 1: Calculate the mean
	var sum float64
	n := float64(len(input))
	for _, value := range input {
		sum += value
	}
	mean := sum / n

	// Step 2: Calculate the sum of squared differences from the mean
	var squaredDiffSum float64
	for _, value := range input {
		diff := value - mean
		squaredDiffSum += math.Pow(diff, 2)
	}

	// Step 3: Calculate the variance
	variance := squaredDiffSum / n // Use n-1 for sample variance


	return variance
}

func Median(numbers []float64) float64 {
	sort.Float64s(numbers)
	n := len(numbers)
	if n%2 == 1 {
		return numbers[n/2]
	}

	return numbers[n/2-1] + numbers[n/2]/2
}