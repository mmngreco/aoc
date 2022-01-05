package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)


func Avg(arr []int) (out int) {
    out = Sum(arr) / len(arr)
    return out
}


func Sum(arr []int) (out int) {
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


func Solve(arr []int) (best int){
    fmt.Println(best)
    var fuel int
    for i, k := range arr {
        fuel = 0
        for _, xi := range arr {
            if xi > k {
                fuel += xi - k
            } else {
                fuel += k - xi
            }
        }
        if fuel < best || i == 0 {
            best = fuel
        }
    }
    return best
}


func main() {
    input := readFile()
    fmt.Println(input)
    fmt.Println("Sum:", Sum(input))
    fmt.Println("Avg:", Avg(input))
    fmt.Println("Solve:", Solve(input))
}

