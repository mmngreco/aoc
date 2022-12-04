package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func part1() {
	b, _ := ioutil.ReadFile(os.Args[1])
	rucksacks := strings.Split(string(b), "\n")
	var sum int
	for _, rucksack := range rucksacks {
		counted := map[rune]bool{}
		n := len(rucksack)
		first := rucksack[:n/2]
		second := rucksack[n/2:]

		for _, item := range first {
			if strings.ContainsRune(second, item) {
				if !counted[item] {
					counted[item] = true
					if (item >= 'a') && (item <= 'z') {
						sum += int(item-'a') + 1
					} else if (item >= 'A') && (item <= 'Z') {
						sum += int(item-'A') + 27
					}
				}
			}
		}
	}

	// Print the sum.
	fmt.Println(sum)
}

const (
	alphabet      = "abcdefghijklmnopqrstuvwxyz"
	upperAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// function to convert rune into int
func convertRuneToInt(r rune) int {
	if (r >= 'a') && (r <= 'z') {
		return int(r-'a') + 1
	} else if (r >= 'A') && (r <= 'Z') {
		return int(r-'A') + 27
	}
	return 0
}

func part2() {
	b, _ := ioutil.ReadFile(os.Args[1])
	rucksacks := strings.Split(string(b), "\n")
	total := 0
	// loop through groups of 3 rucksacks
	for i := 0; i < len(rucksacks)-3; i += 3 {
		// get the 3 rucksacks
		r1 := rucksacks[i]
		r2 := rucksacks[i+1]
		r3 := rucksacks[i+2]
		// loop through the alphabet
		for _, letter := range alphabet + upperAlphabet {
			// check if the letter is in all 3 rucksacks
			if strings.ContainsRune(r1, letter) &&
				strings.ContainsRune(r2, letter) &&
				strings.ContainsRune(r3, letter) {
				// print the letter
				priority := convertRuneToInt(letter)
				// fmt.Println(priority)
				total += priority

			}
		}
		// print a newline
	}
	fmt.Println(total)

}

func main() {
	part2()
}
