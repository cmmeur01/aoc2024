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

	// populate rules and updates
	scanner := bufio.NewScanner(data)
	rules := [][]int{}
	updates := [][]int{}
	spacePassed := false
	for scanner.Scan() {
		line := scanner.Text()
		if spacePassed {
			update := []int{}
			strInts := strings.Split(line, ",")
			for i := 0; i < len(strInts); i++ {
				num, err := strconv.Atoi(strInts[i])
				if err != nil {
					panic(err)
				}
				update = append(update, num)
			}

			updates = append(updates, update)
			continue
		}
		if line == "" {
			spacePassed = true
			continue
		}
		// default add to rules
		strRule := strings.Split(line, "|")
		start, err := strconv.Atoi(strRule[0])
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(strRule[1])
		if err != nil {
			panic(err)
		}

		rules = append(rules, []int{start, end})
	}

	goodUpdates := [][]int{}
	// actually check the updates
	for i := 0; i < len(updates); i++ {
		update := updates[i]
		followsRules := true
		for i := 0; i < len(rules); i++ {
			if !validUpdate(rules[i][0], rules[i][1], update) {
				followsRules = false
				break // go to next update
			}
		}
		if followsRules {
			// been through all the rules this update is good
			goodUpdates = append(goodUpdates, update)
		}
	}

	sum := 0
	for i := 0; i < len(goodUpdates); i++ {
		sum += getMid(goodUpdates[i])
	}

	fmt.Println(sum)
}

func validUpdate(start int, end int, list []int) bool {
	for i := 0; i < len(list); i++ {
		// if the end is in the list we have to check
		// that it is after start if present
		if list[i] == end {
			for j := 0; j < len(list); j++ {
				if list[j] == start {
					return i > j
				}
			}
		}
	}
	return true
}

func getMid(list []int) int {
	len := len(list)
	mid := len / 2
	return list[mid]
}
