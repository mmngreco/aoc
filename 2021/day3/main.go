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


func sum(arr [][]int) ([]int) {
    n := len(arr[0])
    out := make([]int, n)
    for idx, row := range arr {
        for _, elem := range row {
            out[idx] += elem
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


func arr2num(arr []int) (int64){
    str_list := make([]string, len(arr))
    for i, x := range arr {
        str_list[i] = strconv.Itoa(x)
    }

    bin := strings.Join(str_list, "")
    x, _ := strconv.ParseInt(bin, 2, 64)
    return x
}


func gamma(arr2d [][]int) (arr []int) {
    // gamma bits
    arr_t := transpose(arr2d)
    arr_avg := avg(arr_t)
    arr_gt := gt(arr_avg, 0.5)
    return arr_gt
}


func main() {
    file, _ := readFile("input")
    gamma_arr := gamma(file)
    num := arr2num(gamma_arr)
    fmt.Println(gamma_arr)
    fmt.Println(num)
}
