// day 3
package main

import (
    "fmt"
    "strings"
    "io/ioutil"
    "strconv"
)


func readFile(fname string) (array2d [][]int, err error) {

    b, err := ioutil.ReadFile(fname)
    if err != nil { return nil, err }

    lines := strings.Split(strings.Trim(string(b), "\n"), "\n")

    nrows := len(lines)
    ncols := len(lines[0])


    out := make([][]int, nrows)
    row := make([]int, ncols)

    for row_i, l := range lines {
        row2int(l, row)
        out[row_i] = row
    }

    return out, nil
}


func row2int(l string, row []int) {
    for char_i, char := range l {
        // every bit column
        n, err := strconv.Atoi(string(char))
        if err != nil { fmt.Println(err) }
        row[char_i] = n
    }
}


func mode(arr2d [][]int) (arr []int) {

    ncols := len(arr2d[0])
    var counter = make([]map[int]int{0:0;1:0}, ncols)

    for _, row := range arr2d {

        for col_idx, col := range row {

            fmt.Println(counter)
            counter[col_idx].
            counter[col_idx][col] += 1
            fmt.Println(counter)

        }


    }
    return
}


func gammaRate(binary string) (num int, err error) {

    output, err := strconv.ParseInt(binary, 2, 64)

    if err != nil {
        fmt.Println(err)
        return 0, err
    }

    fmt.Printf("Output %d", output)
    return 0, err
}


func main() {
    file, _ := readFile("sample")
    mode(file)
    fmt.Println(file)
    // count
}
