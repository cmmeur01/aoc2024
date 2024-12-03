package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(data)
	safeCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ") // one space between each
		// convert everything in the line to ints could maybe compare strings but meh
		intParts := []int{}
		for i := 0; i < len(parts); i++ {
			val, err := strconv.Atoi(parts[i])
			if err != nil {
				panic(err)
			}
			intParts = append(intParts, val)
		}

		safe := checkLine(intParts)
		if !safe {
			for i := 0; i < len(intParts); i++ {
				// have to make a new array to not jack up the original one
				removed := make([]int, 0, len(intParts)-1)
				removed = append(removed, intParts[:i]...)
				removed = append(removed, intParts[i+1:]...)

				if checkLine(removed) {
					safe = true
					break
				}
			}
		}
		if safe == true {
			safeCount++
		}
	}
	fmt.Println(safeCount)
}

func checkLine(parts []int) bool {
	cur := parts[0]
	next := parts[1]
	safe := false
	if next > cur { // increasing
		safe = increase(parts)
	} else if cur > next { // decreasing
		safe = decrease(parts)
	}

	return safe
}

func increase(parts []int) bool {
	val := parts[0]
	for i := 1; i < len(parts); i++ {
		if val > parts[i] || val == parts[i] || (parts[i]-val) > 3 {
			return false
		}
		val = parts[i]
	}
	return true
}

func decrease(parts []int) bool {
	val := parts[0]
	for i := 1; i < len(parts); i++ {
		if val < parts[i] || val == parts[i] || (val-parts[i]) > 3 {
			return false
		}
		val = parts[i]
	}
	return true
}
