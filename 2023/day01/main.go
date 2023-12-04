package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInputFile() (lines []string) {
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

func part1() {
	lines := readInputFile()
	// fmt.Println(lines)
	numbers := filterNumber(lines)
	// fmt.Println(numbers)
	total := sum(numbers)
	fmt.Println(total)
}

var textNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}

func to_number(word string) int {
	if val, ok := textNumbers[word]; ok {
		return val
	} else {
		number, err := strconv.Atoi(word)
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println("word:", word)
		}
		return number
	}
}

func LineToNumber(line string) int {
	var numberString string
	re := regexp.MustCompile(`(?:one|two|three|four|five|six|seven|eight|nine|zero|\d)`)
	matches := re.FindAllString(line, -1)
	numberString += strconv.Itoa(to_number(matches[0]))
	numberString += strconv.Itoa(to_number(matches[len(matches)-1]))
	number, _ := strconv.Atoi(numberString)
	return number
}

func convertStringsToNumber(lines []string) []int {
	out := make([]int, len(lines))
	for i, line := range lines {
		out[i] = LineToNumber(line)
		// fmt.Println(i, out[i], line)
	}
	return out
}

func part2() {
	lines := readInputFile()
	numbers := convertStringsToNumber(lines)
	fmt.Println(sum(numbers))
}

func main() {
	// part1()
	part2()
}
