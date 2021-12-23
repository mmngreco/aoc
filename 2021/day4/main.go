package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	// "strconv"
	"strings"
)


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
        numbers[i] = v
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
                board_matrix[row][col] = v
                col++
            }
        }
        board_list[ib] = board_matrix
    }
    // fmt.Println(board_list)

    rows_counter := make([][5][]int, len(board_list))
    cols_counter := make([][5][]int, len(board_list))
    // check numbers
    for _, num := range numbers {
        for ib, board := range board_list {
            for ir, row := range board {
                for ic, value := range row {

                    // I should use a long format
                    if num == value {
                        rows_counter[ib][ir] = append(rows_counter[ib][ir], value)
                        cols_counter[ib][ic] = append(cols_counter[ib][ic], value)
                    }

                    if len(cols_counter[ib][ic]) == 5 {
                        break
                        // fmt.Println(rows_counter[ib][ir])
                        // fmt.Println(cols_counter[ib][ic])
                    }
                    if len(rows_counter[ib][ir]) == 5 {
                        break
                        // fmt.Println(rows_counter[ib][ir])
                        // fmt.Println(cols_counter[ib][ic])
                    }
                }
            }
        }
    }
    fmt.Println(rows_counter)

    return file

}


func main() {

    readFile()

}
