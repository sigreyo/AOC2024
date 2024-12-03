package main

import (
	"bufio"
	"fmt"
	"log"
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

	sum := 0

	for i := 0; i < len(left); i++ {
		if slices.Contains(right, left[i]) {
			for _, v := range right {
				if v == left[i] {
					vInt, _ := strconv.Atoi(v)
					sum += vInt
				}
			}
		}
	}

	fmt.Println(sum)
}
