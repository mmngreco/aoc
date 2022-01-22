/*
   v                       v
2 >1< 9  9  9  4  3  2  1 >0<
   ^                       ^
3  9  8  7  8  9  4  9  2  1
      v
9  8 >5< 6  7  8  9  8  9  2
      ^
8  7  6  7  8  9  6  7  8  9
                  v
9  8  9  9  9  6 >5< 6  7  8

*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)


func readFile() {
    // read file
    filename := os.Args[1]
    b, _ := ioutil.ReadFile(filename)
    str := strings.Trim(string(b), "\n")
    lines := strings.Split(str, "\n")

    nrows := len(lines)
    ncols := len(lines[0])

    // Create an empty matrix
    matrix := make([][]uint8, nrows)
    for i := range matrix {
        matrix[i] = make([]uint8, ncols)
    }

    // convert to int
    for row_i, row_v := range lines {
        for col_i, col_v := range row_v {
            v, _ := strconv.Atoi(string(col_v))
            matrix[row_i][col_i] = uint8(v)
        }
    }
    fmt.Println(lines)
    fmt.Println(matrix)

    var is_low_point bool
    // find low points
    for row_i, row_v := range matrix {
        for col_i, col_v := range row_v {
            if (col_v == 9){ continue }
            is_low_point = false

            // TODO we can use a switch
            if (col_i > 0) {
                // left
                is_low_point = is_low_point || matrix[row_i][col_i-1] < col_v
            }
            if (row_i > 0) {
                // down
                is_low_point = is_low_point || matrix[row_i-1][col_i] < col_v
            }
            if (col_i < ncols) {
                // right
                is_low_point = is_low_point || matrix[row_i][col_i+1] < col_v
            }
            if (row_i < nrows) {
                // up
                is_low_point = is_low_point || matrix[row_i+1][col_i] < col_v
            }
            if is_low_point {
                fmt.Println(row_i, col_i, col_v)
            }
        }
    }
    // calculate risk

}


func main() {
    readFile()
}
