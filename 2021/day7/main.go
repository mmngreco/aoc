// Part 1
//
// For example, consider the following horizontal positions:
//
// 16,1,2,0,4,2,7,1,2,14
//
// This means there's a crab with horizontal position 16, a crab with
// horizontal position 1, and so on.
//
// Each change of 1 step in horizontal position of a single crab costs 1 fuel.
// You could choose any horizontal position to align them all on, but the one
// that costs the least fuel is horizontal position 2:
//
// Move from 16 to 2: 14 fuel
// Move from 1 to 2: 1 fuel
// Move from 2 to 2: 0 fuel
// Move from 0 to 2: 2 fuel
// Move from 4 to 2: 2 fuel
// Move from 2 to 2: 0 fuel
// Move from 7 to 2: 5 fuel
// Move from 1 to 2: 1 fuel
// Move from 2 to 2: 0 fuel
// Move from 14 to 2: 12 fuel

// This costs a total of 37 fuel. This is the cheapest possible outcome; more
// expensive outcomes include aligning at position 1 (41 fuel), position 3 (39
// fuel), or position 10 (71 fuel).
//
// Part 2
//
// As it turns out, crab submarine engines don't burn fuel at a constant rate.
// Instead, each change of 1 step in horizontal position costs 1 more unit of
// fuel than the last: the first step costs 1, the second step costs 2, the
// third step costs 3, and so on.
//
// As each crab moves, moving further becomes more expensive. This changes the
// best horizontal position to align them all on; in the example above, this
// becomes 5:
//
// Move from 16 to 5: 66 fuel
// Move from 1 to 5: 10 fuel
// Move from 2 to 5: 6 fuel
// Move from 0 to 5: 15 fuel
// Move from 4 to 5: 1 fuel
// Move from 2 to 5: 6 fuel
// Move from 7 to 5: 3 fuel
// Move from 1 to 5: 10 fuel
// Move from 2 to 5: 6 fuel
// Move from 14 to 5: 45 fuel
//
// This costs a total of 168 fuel. This is the new cheapest possible outcome;
// the old alignment position (2) now costs 206 fuel instead.


package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
    "math"
)


func highestLowest(arr []int) (hi,lo int) {
    hi = math.MaxInt32
    lo = 0

    for i, v := range arr {

        if hi < v  || i == 0 {
            hi = v
        }

        if lo > v || i == 0 {
            lo = v
        }
    }

    return hi, lo
}


func avg(arr []int) (out int) {
    out = sum(arr) / len(arr)
    return out
}


func sum(arr []int) (out int) {

    for _, v := range arr {
        out += v
    }

    return out
}



func readFile() (out []int){

    file := os.Args[1]
    b, _ := ioutil.ReadFile(file)
    input := strings.Trim(string(b), "\n")

    for _, v := range strings.Split(input, ",") {
        vint, _ := strconv.Atoi(v)
        out = append(out, vint)
    }

    return out
}


func distance(arr []int, pos int) (dist int){
    for _, crab_i := range arr {
        dist_i := int(math.Abs(float64(crab_i) - float64(pos)))
        for i := 1; i <= dist_i; i++ {
            dist += i
        }
    }
    return dist
}



func solve(arr []int) (best int){
    var fuelConsumption int

    best = math.MaxInt32
    highest, lowest := highestLowest(arr)
    fmt.Println(lowest, highest)

    for pos := lowest; pos <= highest; pos++ {
        fuelConsumption = distance(arr, pos)
        if fuelConsumption < best {
            best = fuelConsumption
        }
    }

    return best
}


func main() {
    input := readFile()
    // fmt.Println("Sum:", sum(input))
    // fmt.Println("Avg:", avg(input))
    fmt.Println("Solve:", solve(input))
}

