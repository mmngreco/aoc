package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readFile and return lines
func readFile() (lines []string) {
	fname := os.Args[1]
	b, _ := os.ReadFile(fname)
	txt := strings.Trim(string(b), "\n")
	lines = strings.Split(txt, "\n")
	return lines
}

// [asfa2sdf3, a5, as5fas1a32df] -> [23, 5, 52]
func filterNumber(lines []string) (out []int) {
	for _, line := range lines {
		var digits []rune

		for _, char := range line {
			if char >= '0' && char <= '9' {
				digits = append(digits, char)
			}
		}

		// if len(digits) == 1 {
		// 	number, _ := strconv.Atoi(string(digits[0]))
		// 	out = append(out, number)
		// } else if len(digits) > 1 {
		number, _ := strconv.Atoi(string(digits[0]) + string(digits[len(digits)-1]))
		out = append(out, number)
	}
	return out
}

// sum of numbers
func sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func main() {
	lines := readFile()
	// fmt.Println(lines)
	numbers := filterNumber(lines)
	// fmt.Println(numbers)
	total := sum(numbers)
	fmt.Println(total)
}
