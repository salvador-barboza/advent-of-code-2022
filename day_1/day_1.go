package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	scores := make([]int, 0)
	currentElf := 0

	for scanner.Scan() {
		line := scanner.Text()
		calories, _ := strconv.Atoi(line)

		if line == "" {
			scores = append(scores, currentElf)
			currentElf = 0
		} else {
			currentElf += calories
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(scores)))

	fmt.Printf("1) %v, 2) %v, 3) %v\nTotal: %v\n", scores[0], scores[1], scores[2], scores[0]+scores[1]+scores[2])

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
