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

func debug(x ...interface{} ) {
    if DEBUG {
            fmt.Println(x...)
    }
}


func ReadFile(fname string) (array2d [][]int, err error) {

    b, err := ioutil.ReadFile(fname)

    text := string(b)
    lines := strings.Split(strings.Trim(text, "\n"), "\n")
    debug(text)

    nrows := len(lines)
    ncols := len(lines[0])
    out := make([][]int, nrows)

    for row_i, l := range lines {
        row := make([]int, ncols)
        row = Row2int(l, row)
        out[row_i] = row
    }

    return out, nil
}


func Row2int(l string, row []int) ([]int){

    for char_i, char := range l {
        // every bit column
        n, _ := strconv.Atoi(string(char))
        row[char_i] = n
    }
    return row

}


func Transpose(matrix [][]int) (out [][]int) {
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


func Sum2d(matrix [][]int) (out []float32) {

    out = make([]float32, len(matrix))

    for idx, row := range matrix {
        for _, item := range row {
            out[idx] += float32(item)
        }
    }

    return out
}

func Avg2d(matrix [][]int) (out []float32) {

    out = make([]float32, len(matrix))
    n32 := 1 / float32(len(matrix[0]))

    for idx, row := range matrix {
        for _, item := range row {
            out[idx] += float32(item) * n32
        }
    }

    return out
}


func Avg(array []int) (out float32) {
    // Calculate the average of an array.

    n32 := 1 / float32(len(array))

    for _, item := range array {
        out += float32(item) * n32
    }

    return out
}


func GreaterThan(array []float32, num float32) (out []int) {
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


func LessThan(array []float32, num float32) (out []int) {
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



func BitArrayToInt(arr []int) (out int) {
    // Convert bitArray to int

    // Convert int array to str
    str_list := make([]string, len(arr))
    for i, x := range arr {
        str_list[i] = strconv.Itoa(x)
    }
    binary_str := strings.Join(str_list, "")

    // Convert binary str to decimal
    decimal, _ := strconv.ParseInt(binary_str, 2, 64)
    out = int(decimal)

    return out
}


type calculator func([]int) int

func Calculate(matrix [][]int, fn calculator) (int) {
    // oxygen generator rating

    // TODO :
    // - iterate over all rows
    // - keep original position to return the original index of the row
    // - keep relative position after filtering

    // create array of indexes
    var index_list []int
    nrows := len(matrix)
    for irow := 0; irow < nrows; irow++ {
        index_list = append(index_list, irow)
    }

    matrix_copy := matrix

    for icol := 0; icol < nrows; icol++ {
        // cols first
        matrix_copy = Transpose(matrix_copy)
        debug("\n for loop icol: ", icol)

        // get most common bit
        mcb := fn(matrix_copy[icol])
        debug("col:", matrix_copy[icol], "--> mcb: ", mcb)

        // filter file where first bit is equal to the most common bit
        // row first
        matrix_copy = Transpose(matrix_copy)
        matrix_copy = Filter2d(matrix_copy, mcb, icol)
        debug("len:", len(matrix_copy))

        if len(matrix_copy) < 2 { break }

    }
    debug("result", matrix_copy[0])
    return BitArrayToInt(matrix_copy[0])
}


func MostCommonBit(bitColumn []int) int {
    // oxygen generator rating
    avg := Avg(bitColumn)
    if avg >= 0.5 {
        return 1
    } else {
        return 0
    }
}

func LessCommonBit(bitColumn []int) int {
    // oxygen generator rating
    avg := Avg(bitColumn)

    // tricky to see this !!
    if avg < 0.5 {
        return 1
    } else {
        return 0
    }
}

func Epsilon(arr2d [][]int) (out int) {
    // CO2 scrubber rating
    arr_t := Transpose(arr2d)
    arr_avg := Avg2d(arr_t)
    arr_lt := LessThan(arr_avg, 0.5)
    out = BitArrayToInt(arr_lt)

    return out
}


func Filter2d(matrix [][]int, element int, column_look_at int) (out [][]int) {

    matrix = Transpose(matrix)
    debug("# FILTERING COL:", column_look_at, "ELE:", element)
    matrix_column := matrix[column_look_at]
    debug("matrix_column:", matrix_column)

    // find row positions
    var row_positions []int
    for position, value := range matrix_column {
        debug("    val: ", value)
        if value == element {
            debug("        save position: ", position)
            row_positions = append(row_positions, position)
        }
    }

    // select positions
    matrix = Transpose(matrix)
    for _, pos := range row_positions {
        out = append(out, matrix[pos])
    }
    debug("    len: ",len(out))
    return out
}


func main() {
    // read file
    DEBUG = false
    file, _ := ReadFile(os.Args[1])
    fmt.Println(Calculate(file, LessCommonBit))
    fmt.Println(Calculate(file, MostCommonBit) * Calculate(file, LessCommonBit))
}
