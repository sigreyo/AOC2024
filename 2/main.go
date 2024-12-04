package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
	data := make(map[int][]float64, 0)
	i := 0
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		nums := make([]float64, len(line))
		for j, v := range line {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			nums[j] = float64(num)
		}
		data[i] = nums
		i++
	}

	fmt.Println(data)

	safeCount := 0

	for i, line := range data {
		if isSafe(line) {
			safeCount++
			delete(data, i)
		}
	}

	for _, line := range data {
		for j := 0; j < len(line); j++ {
			stripped := append([]float64{}, line[:j]...)
			stripped = append(stripped, line[j+1:]...)
			if isSafe(stripped) {
				safeCount++
				break
			}
		}
	}

	fmt.Println(safeCount)
}

func isSafe(line []float64) bool {
	increasing := true
	decreasing := true

	for i := 0; i < len(line)-1; i++ {

		diff := math.Abs(line[i+1] - line[i])
		if diff < 1 || diff > 3 {
			return false
		}

		if line[i] < line[i+1] {
			decreasing = false
		}

		if line[i] > line[i+1] {
			increasing = false
		}
	}

	return increasing || decreasing
}
