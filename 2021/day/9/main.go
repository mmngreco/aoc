/*

Part 1
------

     2 >1< 9  9  9  4  3  2  1 >0
        ^                       ^
     3  9  8  7  8  9  4  9  2  1
           v
     9  8 >5< 6  7  8  9  8  9  2
           ^
     8  7  6  7  8  9  6  7  8  9
                       v
     9  8  9  9  9  6 >5< 6  7  8

Each number corresponds to the height of a particular location, where 9 is the
highest and 0 is the lowest a location can be.

Your first goal is to find the low points - the locations that are lower than
any of its adjacent locations. Most locations have four adjacent locations (up,
down, left, and right); locations on the edge or corner of the map have three
or two adjacent locations, respectively. (Diagonal locations do not count as
adjacent.)

In the above example, there are four low points, all highlighted: two are in
the first row (a 1 and a 0), one is in the third row (a 5), and one is in the
bottom row (also a 5). All other locations on the heightmap have some lower
adjacent location, and so are not low points.

The risk level of a low point is 1 plus its height. In the above example, the
risk levels of the low points are 2, 1, 6, and 6. The sum of the risk levels of
all low points in the heightmap is therefore 15.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)


const X = 9

var (
    origin = [2]int{0, 0}
    up = [2]int{-1, 0}
    down = [2]int{1, 0}
    left = [2]int{0, -1}
    right = [2]int{0, 1}
)


func is_inside(data [][]int, row_i int, col_i int) (bool) {
    nrows := len(data)
    ncols := len(data[0])
    return ((row_i >= 0) && (row_i < nrows)) && ((col_i >= 0) && (col_i < ncols))
}


func is_basin(data [][]int, row_i int, col_i int) (bool) {
    return is_inside(data, row_i, col_i) && data[row_i][col_i] < X
}


func finder(data [][]int, row_i int, col_i int, motion [2]int, counter *int) int {

    // debug(data)
    // fmt.Printf("row\tcol\n")
    // fmt.Printf("%d+%d\t", row_i, motion[0])
    // fmt.Printf("%d+%d\n", col_i, motion[1])
    row_i = row_i + motion[0]
    col_i = col_i + motion[1]
    // fmt.Println(is_basin(data, row_i, col_i))

    if !is_basin(data, row_i, col_i) {
        return 0
    }

    *counter++
    data[row_i][col_i] = X  // visited

    finder(data, row_i, col_i, up, counter)
    finder(data, row_i, col_i, down, counter)
    finder(data, row_i, col_i, left, counter)
    finder(data, row_i, col_i, right, counter)

    return 1
}


func debug(data [][]int){
    for _, r := range data {
        fmt.Println(r)
    }
}

func make_matrix(nrows int, ncols int) [][]int {
    matrix := make([][]int, nrows)
    for i := range matrix {
        matrix[i] = make([]int, ncols)
    }
    return matrix
}


func fill_matrix(lines []string, matrix [][]int) [][]int {
    for row_i, row_v := range lines {
        for col_i, col_v := range row_v {
            v, _ := strconv.Atoi(string(col_v))
            matrix[row_i][col_i] = int(v)
        }
    }
    return matrix
}


func copy_matrix(matrix [][]int ) [][]int {
    duplicate := make_matrix(len(matrix), len(matrix[0]))
    for i := range matrix {
        duplicate[i] = make([]int, len(matrix[i]))
        copy(duplicate[i], matrix[i])
    }
    return duplicate
}

func is_lowest(matrix [][]int, col_v int, row_i int, col_i int, max_row_i int, max_col_i int) bool {

    if (col_v == 9) {
        return false
    }

    var is_low_point bool
    var left_available, right_available, down_available, up_available bool
    // repeat calculations but it's readable
    up_available = 0 < row_i
    down_available = row_i < max_row_i
    left_available = 0 < col_i
    right_available = col_i < max_col_i

    is_low_point = true

    if left_available {
        // left greater than value
        is_low_point = is_low_point && (matrix[row_i][col_i-1] > col_v)
    }
    if down_available {
        // down greater than value
        is_low_point = is_low_point && (matrix[row_i+1][col_i] > col_v)
    }
    if right_available {
        // right greater than value
        is_low_point = is_low_point && (matrix[row_i][col_i+1] > col_v)
    }
    if up_available {
        // up greater than value
        is_low_point = is_low_point && (matrix[row_i-1][col_i] > col_v)
    }
    return is_low_point
}


func save_largest(largest *[3]int, value int) {
    // FIXME fill all before overwrite a number
    switch {
        case (*largest)[2] < value:
            (*largest)[2] = value
            fmt.Println(2, *largest)
        case (*largest)[1] < value:
            (*largest)[1] = value
            fmt.Println(1, *largest)
        case (*largest)[0] < value:
            (*largest)[0] = value
            fmt.Println(0, *largest)
    }
}


func readFile() {
    // read file
    filename := os.Args[1]
    b, _ := ioutil.ReadFile(filename)
    str := strings.Trim(string(b), "\n")
    lines := strings.Split(str, "\n")

    nrows := len(lines)
    ncols := len(lines[0])

    // Create an empty matrix
    matrix := make_matrix(nrows, ncols)

    // convert to int
    fill_matrix(lines, matrix)

    var is_low_point bool
    var risk int
    var largest [3]int

    max_row_i := nrows - 1
    max_col_i := ncols - 1

    // find low points
    for row_i, row_v := range matrix {
        for col_i, col_v := range row_v {
            is_low_point = is_lowest(matrix, col_v, row_i, col_i, max_row_i, max_col_i)
            if is_low_point {
                // calculate risk
                counter := 0
                data := copy_matrix(matrix)
                finder(data, row_i, col_i, origin, &counter)
                save_largest(&largest, counter)
                risk += 1 + col_v
            }
        }
    }
    fmt.Println(risk, largest)

}


func main() {
    readFile()
}
