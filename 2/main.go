package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	data := make([][]float64, 0)
	for scanner.Scan() {
		trimmed := strings.ReplaceAll(scanner.Text(), " ", ",")
		split := strings.Split(trimmed, ",")

		intLine := make([]float64, 0)
		for _, v := range split {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			intLine = append(intLine, float64(num))
		}

		data = append(data, intLine)
	}

	fmt.Println(data)

	safeCount := 0
	for _, line := range data {
		valid := true

		if !isSorted(line) {
			continue
		}

		for i := 0; i < len(line)-1; i++ {

			diff := math.Abs(line[i] - line[i+1])

			if diff > 3 || diff < 1 {
				valid = false
				break
			}
		}

		if valid {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func isSorted(line []float64) bool {
	if slices.IsSorted(line) {
		fmt.Println(line)
		return true
	}

	slices.Reverse(line)
	return slices.IsSorted(line)
}
