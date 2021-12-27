/*

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

*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)


func readFile() (arr []uint8) {
    file := os.Args[1]
    b, _ := ioutil.ReadFile(file)
    txt := strings.Trim(string(b), "\n")

    for _, v := range strings.Split(txt, ",") {
        val, _ := strconv.Atoi(v)
        arr = append(arr, uint8(val))
    }

    return arr
}



const (
    NewFish uint8 = 8
    Reset uint8 = 6
)


func sim() []uint8 {


    maxDays, _ := strconv.Atoi(os.Args[2])
    arr := readFile()
    // fmt.Println(arr)

    for d := 0; d < maxDays; d++ {
        // iteration: day i
        n := len(arr)
        for i := 0; i < n; i++ {
            if arr[i] == 0 {
                arr[i] = Reset
                arr = append(arr, NewFish)
            } else {
                arr[i]--
            }
        }
        // fmt.Println(arr)
    }
    fmt.Println(len(arr))
    return arr
}


func main() {
    sim()
}
