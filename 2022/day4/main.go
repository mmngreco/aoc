package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1() {
	// Read the list of section assignments from the input.
	reader := bufio.NewReader(os.Stdin)
	var assignments []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		assignments = append(assignments, line)
	}

	// Parse the range of sections for each pair of Elves.
	var ranges [][]int
	for _, a := range assignments {
		a = strings.TrimSpace(a)
		parts := strings.Split(a, ",")
		r1 := strings.Split(parts[0], "-")
		r2 := strings.Split(parts[1], "-")
		ranges = append(ranges, []int{toInt(r1[0]), toInt(r1[1]), toInt(r2[0]), toInt(r2[1])})
	}

	// Count the number of pairs where one range fully contains the other.
	count := 0
	for _, r := range ranges {
		// Check if the first Elf's range fully contains the second Elf's range.
		if (r[0] <= r[2] && r[1] >= r[3]) || (r[0] >= r[2] && r[1] <= r[3]) {
			count++
		}
	}
	fmt.Println("Number of pairs where one range fully contains the other:", count)
}

// toInt converts the given string to an integer.
func toInt(s string) int {
	n := 0
	for _, r := range s {
		n = n*10 + int(r-'0')
	}
	return n
}

func part2() {
	// Read the list of section assignments from the input.
	reader := bufio.NewReader(os.Stdin)
	var assignments []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		assignments = append(assignments, line)
	}

	// Parse the range of sections for each pair of Elves.
	var ranges [][]int
	for _, a := range assignments {
		a = strings.TrimSpace(a)
		parts := strings.Split(a, ",")
		r1 := strings.Split(parts[0], "-")
		r2 := strings.Split(parts[1], "-")
		ranges = append(ranges, []int{toInt(r1[0]), toInt(r1[1]), toInt(r2[0]), toInt(r2[1])})
	}

	// Count the number of pairs that overlap at all.
	count := 0
	for _, r := range ranges {
		// Check if the ranges overlap.
		if (r[0] <= r[3] && r[1] >= r[2]) || (r[0] >= r[3] && r[1] <= r[2]) {
			count++
		}
	}
	fmt.Println("Number of pairs that overlap at all:", count)
}

func main() {
	// part1()
	part2()
}
