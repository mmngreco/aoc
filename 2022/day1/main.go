package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func readFile() (lines []string) {
	fname := os.Args[1]
	b, _ := os.ReadFile(fname)
	txt := strings.Trim(string(b), "\n")
	lines = strings.Split(txt, "\n")
	return lines
}

func sort(arr []int) (out []int) {
	n := len(arr)
	if n < 2 {
		return arr
	}

	idx := rand.Intn(n)
	pivot := arr[idx]

	var (
		left  []int
		right []int
		equal []int
	)
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		}
		if v > pivot {
			right = append(right, v)
		}
		if v == pivot {
			equal = append(equal, v)
		}

	}
	out = append(out, sort(right)...)
	out = append(out, equal...)
	out = append(out, sort(left)...)
	return out
}

func sum(arr []int) (out int) {

	for _, v := range arr {
		out += v
	}
	return out
}

func argmax(groups []int) (imax, max int) {
	for i, e := range groups {
		if e > max {
			imax = i
			max = e
		}
	}
	return imax, max
}

func grouper(lines []string) (groups []int) {
	var total int

	end := len(lines) - 1
	for i, e := range lines {
		lene := len(e)

		if lene > 0 {
			kal, _ := strconv.Atoi(e)
			total += kal
		}

		if (lene == 0) || (i == end) {
			groups = append(groups, total)
			total = 0
		}
	}
	return groups
}

func main1() {
	tst := readFile()
	total := grouper(tst)
	fmt.Println(sort(total))
	fmt.Println(argmax(total))
}

func main() {
	tst := readFile()
	total := grouper(tst)
	top3 := sort(total)[0:3]
	fmt.Println(sum(top3))
}
