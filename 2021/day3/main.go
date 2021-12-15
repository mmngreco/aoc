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


func ReadFile(fname string) (array2d [][]int, err error) {

    b, err := ioutil.ReadFile(fname)

    lines := strings.Split(strings.Trim(string(b), "\n"), "\n")

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


func Gamma(arr2d [][]int) (out int) {
    // oxygen generator rating

    // get most common bit

    bitColumn := 0
    bit := MostCommonBit(file[bitColumn])
    // filter file where first bit is equal to the most common bit
    file := Filter2d(file, bit, bitColumn)
    // repeat with the next bit column until only ramains one entry.
    return out
}


func MostCommonBit(bitColumn []int) int {
    // oxygen generator rating
    avg := Avg(bitColumn)
    if avg > 0.5 {
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


func Filter2d(matrix [][]int, element int, position int) (out [][]int) {
    matrix = Transpose(matrix)
    matrix_column := matrix[position]

    // find row positions
    var row_positions []int
    for position, value := range matrix_column {
        if value == element {
            row_positions = append(row_positions, position)
        }
    }


    // select positions
    for _, pos := range row_positions {
        out = append(out, matrix[pos])
    }
    return out
}




func main() {
    file, _ := ReadFile(os.Args[1])
    // Read the file
    // get most common bit
    bitColumn := 0
    bit := MostCommonBit(file[bitColumn])
    // filter file where first bit is equal to the most common bit
    file := Filter2d(file, bit, bitColumn)
    // repeat with the next bit column until only ramains one entry.
    gamma_rate := Gamma(file)
    epsilon_rate := Epsilon(file)
    fmt.Println(gamma_rate * epsilon_rate)
}
