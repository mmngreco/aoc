// day 3
package main

import (
    "os"
    "fmt"
    "strings"
    "io/ioutil"
    "strconv"
)

var DEBUG bool = true

func debug(x interface{} ) {
    if DEBUG { fmt.Println(x) }
}


func readFile(fname string) (array2d [][]int, err error) {

    b, err := ioutil.ReadFile(fname)

    lines := strings.Split(strings.Trim(string(b), "\n"), "\n")

    nrows := len(lines)
    ncols := len(lines[0])


    out := make([][]int, nrows)

    for row_i, l := range lines {
        row := make([]int, ncols)
        row = row2int(l, row)
        out[row_i] = row
    }

    return out, nil
}


func row2int(l string, row []int) ([]int){

    for char_i, char := range l {
        // every bit column
        n, _ := strconv.Atoi(string(char))
        row[char_i] = n
    }
    return row

}


func transpose(matrix [][]int) (out [][]int) {
    ncols := len(matrix[0])
    nrows := len(matrix)

    for row_i := 0; row_i < ncols; row_i++ {
        row := make([]int, nrows)
        out = append(out, row)
    }

    for row_i := 0; row_i < nrows; row_i++ {
        for col_i := 0; col_i < ncols; col_i++ {
            out[col_i][row_i] = matrix[row_i][col_i]
        }
    }

    return out
}


func sum(matrix [][]int) (out []float32) {

    out = make([]float32, len(matrix))

    for idx, row := range matrix {
        for _, item := range row {
            out[idx] += float32(item)
        }
    }

    return out
}

func avg(matrix [][]int) (out []float32) {

    out = make([]float32, len(matrix))
    n32 := 1 / float32(len(matrix[0]))

    for idx, row := range matrix {
        for _, item := range row {
            out[idx] += float32(item) * n32
        }
    }

    return out
}


func gt(array []float32, num float32) (out []int) {
    n := len(array)
    out = make([]int, n)

    for idx, item := range array {
        if item > num {
            out[idx] = 1
        } else if item < num {
            out[idx] = 0
        }
    }

    return out
}


func lt(array []float32, num float32) (out []int) {
    n := len(array)
    out = make([]int, n)

    for idx, item := range array {
        if item < num {
            out[idx] = 1
        } else if item > num {
            out[idx] = 0
        }
    }

    return out
}



func arr2num(arr []int) (out int){
    str_list := make([]string, len(arr))
    for i, x := range arr {
        str_list[i] = strconv.Itoa(x)
    }

    bin := strings.Join(str_list, "")
    x, _ := strconv.ParseInt(bin, 2, 64)
    out = int(x)
    return out
}


func gamma(arr2d [][]int) (out int) {
    // gamma
    arr_t := transpose(arr2d)
    arr_avg := avg(arr_t)
    arr_gt := gt(arr_avg, 0.5)
    out = arr2num(arr_gt)
    return out
}


func epsilon(arr2d [][]int) (out int) {
    // epsilon
    arr_t := transpose(arr2d)
    arr_avg := avg(arr_t)
    arr_lt := lt(arr_avg, 0.5)
    out = arr2num(arr_lt)
    return out
}



func main() {
    file, _ := readFile(os.Args[1])
    gamma_rate := gamma(file)
    epsilon_rate := epsilon(file)
    fmt.Println(gamma_rate * epsilon_rate)
}
