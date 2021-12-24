/* Day 4: Bingo!

7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7


The score of the winning board can now be calculated. Start by finding the sum
of all unmarked numbers on that board; in this case, the sum is 188. Then,
multiply that sum by the number that was just called when the board won, 24, to
get the final score, 188 * 24 = 4512.

*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	// "strconv"
	"strings"
)


func Sum2d(arr [5][5]int) int {
    var out int = 0
    for _, row := range arr {
        for _, ele := range row {
            if ele == 0 { continue }
            out += ele - 1
        }
    }
    return out
}


// Read files from arg
func readFile() string {
    // read file from arguments
    input := os.Args[1]
    b, _ := ioutil.ReadFile(input)
    file := string(b)


    // array of numbers to check in each board
    guest_arr := strings.Split(strings.Split(file, "\n\n")[0], ",")
    numbers := make([]int, len(guest_arr))
    for i, num := range guest_arr {
        v, _ := strconv.Atoi(num)
        numbers[i] = v + 1
    }

    // each board
    board_str_list := strings.Split(file, "\n\n")[1:]
    board_list := make([][5][5]int, len(board_str_list))
    for ib, board := range board_str_list {
        board := strings.Split(board, "\n")
        var board_matrix [5][5]int
        var col int
        for row, row_str := range board {
            col = 0
            for _, e := range strings.Split(row_str, " "){
                if len(e) == 0 { continue }
                v, _ := strconv.Atoi(e)
                board_matrix[row][col] = v + 1
                col++
            }
        }
        board_list[ib] = board_matrix
    }
    // fmt.Println(board_list)


    // check numbers
    nboard := len(board_list)
    rows_counter := make([][5][]int, nboard)
    cols_counter := make([][5][]int, nboard)

    num := 0
    value := 0
    done := 0
    points := 0

    for inum:=0; (done == 0) && (inum < len(numbers)); inum++ {

        num = numbers[inum]

        for iboard:=0; (iboard < nboard) && (done == 0); iboard++ {

            board := &board_list[iboard]
            // fmt.Println(iboard, *board)

            for irow:=0; (irow < 5) && (done == 0); irow++{

                row := &board[irow]
                // fmt.Println(iboard, irow, *row)

                for icol:=0; (icol < 5) && (done == 0); icol++ {
                    value = (*row)[icol]
                    row_current := &rows_counter[iboard][irow]
                    col_current := &cols_counter[iboard][icol]

                    if num == value {
                        *row_current = append((*row_current), num-1)
                        *col_current = append((*col_current), num-1)
                        (*row)[icol] = 0
                    }

                    if len(*row_current) == 5 {
                        done = 1
                        points = Sum2d(*board) * (*row_current)[4]
                    }

                    if len(cols_counter[iboard][icol]) == 5 {
                        done = 1
                        points = Sum2d(*board)
                    }
                }
            }
        }
    }
    fmt.Println(points)

    return file

}


func main() {

    readFile()

}
