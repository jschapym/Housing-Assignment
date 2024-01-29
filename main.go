package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	N := 100
	file, err := os.Create("housesOutputGo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < N; i++ {
		houses, err := readCSV("housesInput.csv")
		if err != nil {
			log.Fatal(err)
		}
		stats := describe(houses)
		for _, line := range stats {
			fmt.Fprintln(file, line)
		}
		fmt.Fprintln(file, "")
	}
}

func readCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ',' // Set the delimiter if it's not a comma
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func describe(data [][]string) []string {
	numRows := len(data)
	if numRows == 0 {
		return []string{"No data to describe"}
	}

	numCols := len(data[0])
	stats := make([]string, numCols)

	for col := 0; col < numCols; col++ {
		var values []float64
		sum := 0.0

		for row := 1; row < numRows; row++ {
			value, err := strconv.ParseFloat(data[row][col], 64)
			if err == nil {
				values = append(values, value)
				sum += value
			}
		}

		if len(values) == 0 {
			stats[col] = fmt.Sprintf("Column %d: No numeric data", col+1)
			continue
		}

		mean := sum / float64(len(values))
		min, max := findMinMax(values)
		stdDev := calculateStdDev(values, mean)

		stats[col] = fmt.Sprintf("Column %d: Mean=%.2f, Min=%.2f, Max=%.2f, StdDev=%.2f", col+1, mean, min, max, stdDev)
	}

	return stats
}

func findMinMax(values []float64) (min, max float64) {
	if len(values) == 0 {
		return 0, 0
	}

	min, max = values[0], values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func calculateStdDev(values []float64, mean float64) float64 {
	if len(values) == 0 {
		return 0
	}

	var sumSqDiff float64
	for _, v := range values {
		diff := v - mean
		sumSqDiff += diff * diff
	}

	variance := sumSqDiff / float64(len(values))
	stdDev := math.Sqrt(variance)
	return stdDev
}
