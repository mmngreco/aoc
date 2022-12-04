package main

// first column
// ------------
// A for Rock
// B for Paper
// C for Scissors.

// second column
// -------------
// X for Rock
// Y for Paper
// Z for Scissors.

// points
// ------
// 1 for Rock
// 2 for Paper
// 3 for Scissors

// game points
// -----------
// 0 if I lose
// 3 if draw
// 6 if I win

// test
// ----
// A Y  -> 1 2 ->  I win  -> 0 + 2 + 6 -> 8
// B X  -> 2 1 ->  I lose -> 0 + 1 + 0 -> 1
// C Z  -> 3 3 ->  Tie    -> 0 + 3 + 3 -> 6
// result : 15

// algorithm
// input -> points -> result -> score -> sum

// What would your total score be if everything goes exactly according to your
// strategy guide?

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	rock     int = 1
	paper    int = 2
	scissors int = 3
)

var choice2points = map[string]int{
	// opponent
	"A": rock,
	"B": paper,
	"C": scissors,
	// me
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

var game_points = map[string]int{
	"draw": 3,
	"win":  6,
	"lose": 0,
}

func resolver(op, my int) (result string) {
	if op == my {
		result = "draw"
	}

	if op == rock {
		if my == paper {
			result = "win"
		}
		if my == scissors {
			result = "lose"
		}
	}

	if op == paper {
		if my == rock {
			result = "lose"
		}
		if my == scissors {
			result = "win"
		}
	}

	if op == scissors {
		if my == rock {
			result = "win"
		}
		if my == paper {
			result = "lose"
		}
	}
	return result
}

func game(op, my string) (out int) {
	op_pts := choice2points[op]
	my_pts := choice2points[my]
	myresult := resolver(op_pts, my_pts)
	res := game_points[myresult]
	out = my_pts + res

	return out
}

func readfile() (lines []string) {
	file := os.Args[1]
	fmt.Println(file)
	b, _ := ioutil.ReadFile(file)
	txt := string(b)
	txt = strings.Trim(txt, "\n")
	lines = strings.Split(txt, "\n")
	return lines
}

func part1() {
	lines := readfile()
	total := 0
	for _, v := range lines {
		round := strings.Split(v, " ")
		// fmt.Printf("%q\n", round)
		res := game(round[0], round[1])
		// fmt.Printf("%d\n", res)
		total += res
	}
	fmt.Printf("part 1: %d\n", total)
}

// part 2
// X means you need to lose
// Y means you need to end the round in a draw
// Z means you need to win

var choice2result = map[string]string{
	"X": "lose",
	"Y": "draw",
	"Z": "win",
}

func solver(op, res string) (my string) {
	op_pts := choice2points[op]
	options := []string{"X", "Y", "Z"}
	for _, v := range options {
		my_pts := choice2points[v]
		myresult := resolver(op_pts, my_pts)
		if myresult == res {
			my = v
		}
	}
	return my
}

func part2() {
	lines := readfile()
	total := 0
	for _, v := range lines {
		round := strings.Split(v, " ")
		op := round[0]
		res := round[1]
		// fmt.Printf("%q\n", round)
		expected := choice2result[res]
		my := solver(op, expected)
		// fmt.Printf("expected %s\nmust play %s\n", res, my)
		res_pts := game(op, my)
		// fmt.Printf("%d\n", res_pts)
		total += res_pts
	}
	// fmt.Println(total)
	fmt.Printf("part 2: %d\n", total)
}

func main() {
	part2()
}
