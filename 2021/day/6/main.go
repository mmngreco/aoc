/*

--- Part One ---
So, suppose you have a lanternfish with an internal timer value of 3:

- After one day, its internal timer would become 2.
- After another day, its internal timer would become 1.
- After another day, its internal timer would become 0.
- After another day, its internal timer would reset to 6, and it would create a
  new lanternfish with an internal timer of 8.
- After another day, the first lanternfish would have an internal timer of 5,
  and the second lanternfish would have an internal timer of 7.

Initial state: 3,4,3,1,2
After  1 day:  2,3,2,0,1
After  2 days: 1,2,1,6,0,8
After  3 days: 0,1,0,5,6,7,8
After  4 days: 6,0,6,4,5,6,7,8,8
After  5 days: 5,6,5,3,4,5,6,7,7,8
After  6 days: 4,5,4,2,3,4,5,6,6,7
After  7 days: 3,4,3,1,2,3,4,5,5,6
After  8 days: 2,3,2,0,1,2,3,4,4,5
After  9 days: 1,2,1,6,0,1,2,3,3,4,8
After 10 days: 0,1,0,5,6,0,1,2,2,3,7,8
After 11 days: 6,0,6,4,5,6,0,1,1,2,6,7,8,8,8
After 12 days: 5,6,5,3,4,5,6,0,0,1,5,6,7,7,7,8,8
After 13 days: 4,5,4,2,3,4,5,6,6,0,4,5,6,6,6,7,7,8,8
After 14 days: 3,4,3,1,2,3,4,5,5,6,3,4,5,5,5,6,6,7,7,8
After 15 days: 2,3,2,0,1,2,3,4,4,5,2,3,4,4,4,5,5,6,6,7
After 16 days: 1,2,1,6,0,1,2,3,3,4,1,2,3,3,3,4,4,5,5,6,8
After 17 days: 0,1,0,5,6,0,1,2,2,3,0,1,2,2,2,3,3,4,4,5,7,8
After 18 days: 6,0,6,4,5,6,0,1,1,2,6,0,1,1,1,2,2,3,3,4,6,7,8,8,8,8

--- Part Two ---
Suppose the lanternfish live forever and have unlimited food and space. Would they take over the entire ocean?

After 256 days in the example above, there would be a total of 26984457539 lanternfish!

How many lanternfish would there be after 256 days?
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)


func readFile() (arr []uint64) {
    file := os.Args[1]
    b, _ := ioutil.ReadFile(file)
    txt := strings.Trim(string(b), "\n")

    for _, v := range strings.Split(txt, ",") {
        val, _ := strconv.Atoi(v)
        arr = append(arr, uint64(val))
    }

    return arr
}



const (
    NewFish uint64 = 8
    Reset uint64 = 6
)


func sim() []uint64 {


    maxDays, _ := strconv.Atoi(os.Args[2])
    arr := readFile()

    lifeCycle := [9]uint64{}
    for _, life := range arr {
        if life == 0 { lifeCycle[0]++ }
        if life == 1 { lifeCycle[1]++ }
        if life == 2 { lifeCycle[2]++ }
        if life == 3 { lifeCycle[3]++ }
        if life == 4 { lifeCycle[4]++ }
        if life == 5 { lifeCycle[5]++ }
        if life == 6 { lifeCycle[6]++ }
    }

    for d := 1; d <= maxDays; d++ {
        for i, n := range lifeCycle {
            // i = 0, n = 0
            if i == 1 { lifeCycle[1] -= n; lifeCycle[0] += n }
            if i == 2 { lifeCycle[2] -= n; lifeCycle[1] += n }
            if i == 3 { lifeCycle[3] -= n; lifeCycle[2] += n }
            if i == 4 { lifeCycle[4] -= n; lifeCycle[3] += n }
            if i == 5 { lifeCycle[5] -= n; lifeCycle[4] += n }
            if i == 6 { lifeCycle[6] -= n; lifeCycle[5] += n }
            if i == 7 { lifeCycle[7] -= n; lifeCycle[6] += n }
            if i == 8 { lifeCycle[8] -= n; lifeCycle[7] += n }
            if i == 0 { lifeCycle[0] -= n; lifeCycle[6] += n  ; lifeCycle[8] += n }
        }
        // fmt.Println(lifeCycle)
        // fmt.Println(Sum(lifeCycle))

    }
    // fmt.Println(lifeCycle)
    fmt.Println(Sum(lifeCycle))
    return lifeCycle[:]
}


func Sum(arr [9]uint64) (out uint64) {
    for _, v := range arr {
        out += v
    }
    return out
}


func main() {
    sim()
}
