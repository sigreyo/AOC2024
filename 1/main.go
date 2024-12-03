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

	left := []string{}
	right := []string{}

	for scanner.Scan() {
		trimmed := strings.Split(scanner.Text(), "   ")

		left = append(left, trimmed[0])
		right = append(right, trimmed[1])
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0

	for i := 0; i < len(left); i++ {
		leftInt, _ := strconv.Atoi(left[i])
		rightInt, _ := strconv.Atoi(right[i])

		sum += int(math.Abs(float64(leftInt - rightInt)))
	}

	fmt.Println(sum)

}
