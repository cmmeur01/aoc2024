package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	// populate matrix
	matrix := [][]string{}
	scanner := bufio.NewScanner(data)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		ltrs := strings.Split(line, "")
		matrix = append(matrix, []string{})
		matrix[i] = append(matrix[i], ltrs...)
		i++
	}

	count := 0
	dirs := [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// If we find an X, check for adjacent XMAS pattern
			if matrix[i][j] == "X" {
				m := getLetterPos(matrix, []int{i, j}, dirs, "M")
				if len(m) > 0 {
					a := getLetterPos(matrix, m, dirs, "A")
					if len(a) > 0 {
						s := getLetterPos(matrix, a, dirs, "S")
						if len(s) > 0 {
							count++
						}
					}
				}
			}
		}
	}

	fmt.Println(count)
}

func isValidPos(matrix [][]string, x int, y int) bool {
	return x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0])
}

func getLetterPos(matrix [][]string, start []int, dirs [][]int, target string) []int {
	for _, dir := range dirs {
		newPos := []int{start[0] + dir[0], start[1] + dir[1]}
		if isValidPos(matrix, newPos[0], newPos[1]) {
			if matrix[newPos[0]][newPos[1]] == target {
				return newPos
			}
		}
	}
	return []int{}
}
