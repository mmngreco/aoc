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
    fmt.Println("Sum:", sum(input))
    fmt.Println("Avg:", avg(input))
    fmt.Println("Solve:", solve(input))
}

