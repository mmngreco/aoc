package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
    // "reflect"
)

// It would be better for such a function to return error, instead of handling
// it on their own.
func readFile(fname string) (nums []int, err error) {

    b, err := ioutil.ReadFile(fname)
    if err != nil { return nil, err }

    lines := strings.Split(string(b), "\n")
    nums = make([]int, 0, len(lines))

    for _, l := range lines {

        // Empty line occurs at the end of the file when we use Split.
        if len(l) == 0 { continue }

        // Atoi better suits the job when we know exactly what we're dealing
        // with. Scanf is the more general option.
        n, err := strconv.Atoi(l)
        if err != nil { return nil, err }
        nums = append(nums, n)
    }

    return nums, nil
}


type convert2 func(int, int) int
func diff2(a int, b int) int { return a - b }

type convert func(int) int
func count_pos(a int) int { if a > 0 {return 1 } else {return 0}}
func sum(a int, b int) int { return a + b }


func apply2(arr1 []int, arr2 []int, fn convert2) []int {
    // Apply a function over each element in both arr1 and arr2.

    var out []int

    n := len(arr1)

    for i := 0; i < n; i++ {
        out_i := fn(arr1[i], arr2[i])
        out = append(out, out_i)
    }

    return out
}


func apply(arr1 []int, fn convert) []int {
    // Apply a function over a single array.

    var out []int

    n := len(arr1)

    for i := 0; i < n; i++ {
        out_i := fn(arr1[i])
        out = append(out, out_i)
    }

    return out
}


func reduce(arr []int, fn convert2) int {

    var out int = arr[0]

    n := len(arr)

    for i := 1; i < n; i++ {
        current := arr[i]
        out = fn(out, current)
    }

    return out

}


func main() {
    // read file
    // arr, _ := readFile("day1.sample")
    arr, _ := readFile("day1.input")

    // params
    n := len(arr)
    win := 3

    // get increments
    increments := apply2(arr[win:n], arr[0:n-win], diff2)

    // count positives
    positives := apply(increments, count_pos)
    count := reduce(positives, sum)

    // fmt.Println("increments:", increments)
    // fmt.Println("positives:", positives)
    fmt.Println("count:", count)

}
