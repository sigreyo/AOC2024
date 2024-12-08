package main

import (
	"bufio"
	"fmt"
	"log"
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

	sum := 0
	for scanner.Scan() {
		muls := extractMuls(scanner.Text())
		sum += splitAndMultiPlyMuls(muls)
	}

	fmt.Println(sum)
}

func extractMuls(line string) []string {
	muls := make([]string, 0)
	toAdd := ""
	for i, v := range line {
		if i > 3 && v == '(' && line[i-3:i] == "mul" && toAdd == "" {
			toAdd = string(line[i])
			continue
		} else if strings.Contains("0123456789,", string(v)) && toAdd != "" {
			toAdd += string(v)
		} else if string(v) == ")" && toAdd != "" {
			toAdd += string(v)
			muls = append(muls, toAdd[1:len(toAdd)-1])
			toAdd = ""
		} else {
			toAdd = ""
		}
	}

	return muls
}

func splitAndMultiPlyMuls(muls []string) int {
	sum := 0
	for _, v := range muls {
		if len(v) < 3 {
			continue
		}

		values := strings.Split(v, ",")

		left, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatal(err)
		}
		right, err := strconv.Atoi(values[1])
		if err != nil {
			log.Fatal(err)
		}

		sum += left * right
	}

	return sum
}
