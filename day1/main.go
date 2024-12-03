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

	left := []int{}
	right := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ") // three spaces between each
		first, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		left = append(left, first)
		second, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		right = append(right, second)
	}

	if len(left) != len(right) {
		fmt.Println("arrays are not the same length")
	}

	// part 1
	// sort.Ints(left)
	// sort.Ints(right)
	// totalDistance := 0
	// for i := 0; i < len(left); i++ {
	// 	distance := left[i] - right[i]
	// 	if distance < 0 {
	// 		distance = -distance
	// 	}
	// 	totalDistance += distance
	// }
	// fmt.Println(totalDistance)

	// part 2
	rightMap := make(map[int]int)
	for i := 0; i < len(right); i++ {
		val := right[i]
		cur := rightMap[val]
		rightMap[val] = cur + 1
	}

	simTotal := 0
	for i := 0; i < len(left); i++ {
		val := left[i]
		numRightTimes := rightMap[val]

		simTotal += (val * numRightTimes)
	}
	fmt.Println(simTotal)
}
