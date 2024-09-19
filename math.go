package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <file path>", os.Args[0])
	}

	filePath := os.Args[1]
	numbers, err := readFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	if len(numbers) == 0 {
		log.Fatal("No numbers in file")
	}

	average := calculateAverage(numbers)
	median := calculateMedian(numbers)
	variance := calculateVariance(numbers, average)
	stdDev := calculateStdDev(variance)

	fmt.Printf("Average: %.2f\n", average)
	fmt.Printf("Median: %.2f\n", median)
	fmt.Printf("Variance: %.2f\n", variance)
	fmt.Printf("Standard Deviation: %.2f\n", stdDev)
}

func readFile(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return numbers, nil
}

func calculateAverage(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func calculateMedian(numbers []float64) float64 {
	sort.Float64s(numbers)
	mid := len(numbers) / 2
	if len(numbers)%2 == 0 {
		return (numbers[mid-1] + numbers[mid]) / 2
	}
	return numbers[mid]
}

func calculateVariance(numbers []float64, mean float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += (num - mean) * (num - mean)
	}
	return sum / float64(len(numbers))
}

func calculateStdDev(variance float64) float64 {
	return math.Sqrt(variance)
}
