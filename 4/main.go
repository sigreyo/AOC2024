package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	matrix := make([]string, 0)
	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}
	// count := 0

	// count += findNormalAndBackwardsXMAS(matrix)
	// count += findHorizontalXMAS(matrix)
	// count += findDiagonalXMAS(matrix)
	count := diamondSearch(matrix)

	fmt.Println(count)
}

func diamondSearch(matrix []string) int {
	count := 0
	for matrixIndex, line := range matrix {
		for i := range line {
			if string(line[i]) == "X" {

				//normal and backwards
				if i+3 < len(line) {
					word := string(line[i+1]) + string(line[i+2]) + string(line[i+3])
					if word == "MAS" {
						count++
					}
				}
				if i-3 >= 0 {
					word := string(line[i-1]) + string(line[i-2]) + string(line[i-3])
					if word == "MAS" {
						count++
					}
				}

				//down
				if matrixIndex+3 < len(matrix) {
					word := string(matrix[matrixIndex+1][i]) + string(matrix[matrixIndex+2][i]) + string(matrix[matrixIndex+3][i])
					if word == "MAS" {
						count++
					}
				}

				//up
				if matrixIndex-3 >= 0 {
					word := string(matrix[matrixIndex-1][i]) + string(matrix[matrixIndex-2][i]) + string(matrix[matrixIndex-3][i])
					if word == "MAS" {
						count++
					}
				}

				//down right
				if matrixIndex+3 < len(matrix) && i+3 < len(line) {
					word := string(matrix[matrixIndex+1][i+1]) + string(matrix[matrixIndex+2][i+2]) + string(matrix[matrixIndex+3][i+3])
					if word == "MAS" {
						count++
					}
				}

				//down left
				if matrixIndex+3 < len(matrix) && i-3 >= 0 {
					word := string(matrix[matrixIndex+1][i-1]) + string(matrix[matrixIndex+2][i-2]) + string(matrix[matrixIndex+3][i-3])
					if word == "MAS" {
						count++
					}
				}

				//up right
				if matrixIndex-3 >= 0 && i+3 < len(line) {
					word := string(matrix[matrixIndex-1][i+1]) + string(matrix[matrixIndex-2][i+2]) + string(matrix[matrixIndex-3][i+3])
					if word == "MAS" {
						count++
					}
				}

				//up left
				if matrixIndex-3 >= 0 && i-3 >= 0 {
					word := string(matrix[matrixIndex-1][i-1]) + string(matrix[matrixIndex-2][i-2]) + string(matrix[matrixIndex-3][i-3])
					if word == "MAS" {
						count++
					}

				}
			}
		}
	}

	return count
}

// func findNormalAndBackwardsXMAS(matrix []string) int {
// 	count := 0
// 	for _, line := range matrix {
// 		for i := range line {
// 			if i+4 < len(line) && string(line[i:i+4]) == "XMAS" {
// 				count++
// 			}
// 			if i-4 >= 0 && string(line[i-4:i]) == "SAMX" {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }

// func findHorizontalXMAS(matrix []string) int {
// 	count := 0
// 	for matrixIndex, line := range matrix {
// 		for i := range line {
// 			if string(line[i]) == "X" {
// 				if matrixIndex+3 < len(matrix) {
// 					word := strings.Join([]string{
// 						string(line[i]),
// 						string(matrix[matrixIndex+1][i]),
// 						string(matrix[matrixIndex+2][i]),
// 						string(matrix[matrixIndex+3][i]),
// 					}, "")
// 					if word == "XMAS" || word == "SAMX" {
// 						count++
// 					}
// 				}
// 				if matrixIndex-3 >= 0 {
// 					word := strings.Join([]string{
// 						string(line[i]),
// 						string(matrix[matrixIndex-1][i]),
// 						string(matrix[matrixIndex-2][i]),
// 						string(matrix[matrixIndex-3][i]),
// 					}, "")
// 					if word == "XMAS" || word == "SAMX" {
// 						count++
// 					}
// 				}
// 			}
// 		}

// 	}
// 	return count
// }

// func findDiagonalXMAS(matrix []string) int {
// 	count := 0
// 	for matrixIndex, line := range matrix {
// 		for i := range line {
// 			if string(line[i]) == "X" {
// 				//down right
// 				if matrixIndex+3 < len(matrix) && i+3 < len(line) {
// 					word := strings.Join([]string{
// 						string(line[i]),
// 						string(matrix[matrixIndex+1][i+1]),
// 						string(matrix[matrixIndex+2][i+2]),
// 						string(matrix[matrixIndex+3][i+3]),
// 					}, "")
// 					if word == "XMAS" || word == "SAMX" {
// 						count++
// 					}
// 				}

// 				//down left
// 				if matrixIndex+3 < len(matrix) && i-3 >= 0 {
// 					word := strings.Join([]string{
// 						string(line[i]),
// 						string(matrix[matrixIndex+1][i-1]),
// 						string(matrix[matrixIndex+2][i-2]),
// 						string(matrix[matrixIndex+3][i-3]),
// 					}, "")
// 					if word == "XMAS" || word == "SAMX" {
// 						count++
// 					}
// 				}

// 				//up right
// 				if matrixIndex-3 >= 0 && i+3 < len(line) {
// 					word := strings.Join([]string{
// 						string(line[i]),
// 						string(matrix[matrixIndex-1][i+1]),
// 						string(matrix[matrixIndex-2][i+2]),
// 						string(matrix[matrixIndex-3][i+3]),
// 					}, "")
// 					if word == "XMAS" || word == "SAMX" {
// 						count++
// 					}
// 				}

// 				//up left
// 				if matrixIndex-3 >= 0 && i-3 >= 0 {
// 					word := strings.Join([]string{
// 						string(line[i]),
// 						string(matrix[matrixIndex-1][i-1]),
// 						string(matrix[matrixIndex-2][i-2]),
// 						string(matrix[matrixIndex-3][i-3]),
// 					}, "")
// 					if word == "XMAS" || word == "SAMX" {
// 						count++
// 					}
// 				}

// 			}
// 		}
// 	}
// 	return count
// }
