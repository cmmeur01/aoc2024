package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	pattern := `mul\((\d{1,3}),\s*(\d{1,3})\)`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("regex is broke")
		panic(err)
	}

	sum := 0
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindAllString(line, -1)

		sum += processMatches(matches)
	}

	fmt.Println(sum)
}

func processMatches(matches []string) int {
	sum := 0
	for i := 0; i < len(matches); i++ {
		strMatch := matches[i]
		nums := strings.TrimPrefix(strMatch, "mul(")
		nums = strings.TrimSuffix(nums, ")")
		strInts := strings.Split(nums, ",")

		sum += multiplyStringInts(strInts)

	}
	return sum
}

func multiplyStringInts(intString []string) int {
	one, err := strconv.Atoi(intString[0])
	if err != nil {
		panic(err)
	}

	two, err := strconv.Atoi(intString[1])
	if err != nil {
		panic(err)
	}

	return one * two
}
