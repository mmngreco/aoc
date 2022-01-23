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


func readFile() {
    // read file
    filename := os.Args[1]
    b, _ := ioutil.ReadFile(filename)
    str := strings.Trim(string(b), "\n")
    lines := strings.Split(str, "\n")

    nrows := len(lines)
    ncols := len(lines[0])

    // Create an empty matrix
    matrix := make([][]int, nrows)
    for i := range matrix {
        matrix[i] = make([]int, ncols)
    }

    // convert to int
    for row_i, row_v := range lines {
        for col_i, col_v := range row_v {
            v, _ := strconv.Atoi(string(col_v))
            matrix[row_i][col_i] = int(v)
        }
    }

    var is_low_point bool
    var risk int
    var left_available, right_available, down_available, up_available bool

    max_row_i := nrows - 1
    max_col_i := ncols - 1
    // find low points
    for row_i, row_v := range matrix {

        down_available = row_i > 0
        up_available = row_i < max_row_i
        left_available = false
        right_available = true

        for col_i, col_v := range row_v {

            // early stop
            if (col_v == 9){ continue }

            is_low_point = true
            right_available = (col_i < max_col_i)

            if left_available {
                // left
                is_low_point = is_low_point && (matrix[row_i][col_i-1] > col_v)
            }
            if down_available {
                // down
                is_low_point = is_low_point && (matrix[row_i-1][col_i] > col_v)
            }
            if right_available {
                // right
                is_low_point = is_low_point && (matrix[row_i][col_i+1] > col_v)
            }
            if up_available {
                // up
                is_low_point = is_low_point && (matrix[row_i+1][col_i] > col_v)
            }

            if is_low_point {
                // calculate risk
                risk += 1 + col_v
            }

            left_available = true

        }
    }
    fmt.Println(risk)

}



func main() {
    readFile()
}
