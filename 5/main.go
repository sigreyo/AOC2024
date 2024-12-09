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

	addUpdates := false
	rules := make([][]string, 0)
	updates := make([]string, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			addUpdates = true
			continue
		}
		if !addUpdates {
			rules = append(rules, strings.Split(scanner.Text(), "|"))
		} else {
			updates = append(updates, scanner.Text())
		}
	}

	sum := 0
	for _, v := range updates {
		matchingRules := findMatchingRules(rules, v)
		sum += checkUpdate(v, matchingRules)
	}

	fmt.Println(sum)
}

func findMatchingRules(rules [][]string, update string) [][]string {
	matchingRules := make([][]string, 0)
	for _, rule := range rules {
		if strings.Contains(update, rule[0]) && strings.Contains(update, rule[1]) {
			matchingRules = append(matchingRules, rule)
		}
	}
	return matchingRules
}

func checkUpdate(update string, rules [][]string) int {
	splitted := strings.Split(update, ",")
	valid := false
	for _, rule := range rules {
		if slices.Index(splitted, rule[0]) < slices.Index(splitted, rule[1]) {
			valid = true
		} else {
			valid = false
			break
		}
	}
	if valid {
		num, _ := strconv.Atoi(splitted[len(splitted)/2])
		return num
	}

	return 0
}
