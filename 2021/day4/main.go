package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)


type Board struct {
    dim int
    raw string
    row_list []string
    // attributes
    rows [][]int
    cols [][]int
    done bool
    iters int
    last string
    win_index []int
    win_values []string
}

// Set raw and list
func (self *Board) set_board(board string, dim int) {

    self.iters = -1
    self.raw = board
    self.row_list = self.clean_raw(board)
    self.init(dim)

}

// Check number
func (self *Board) mark_number(number string) bool {
    self.iters++
    list := self.row_list

    for ielem, elem := range list {

        if number == elem {

            self.append_row(ielem)
            self.append_col(ielem)
            self.last = elem
            if self.done {
                return self.done
            }

        }
    }

    return self.done
}

func (self *Board) append_row(index int) bool {
    // If dim = 3:
    //    e : i , j
    //   ---:---,--
    //    0 : 0 , 0
    //    1 : 0 , 1
    //    2 : 0 , 2
    //    3 : 1 , 0
    //    4 : 1 , 1
    //    5 : 1 , 2

    var i int
    i = index % self.dim
    self.rows[i] = append(self.rows[i], index)

    if len(self.rows[i]) == self.dim {
        self.done = true
        self.win_index = self.rows[i]
    }
    return self.done

}

func (self *Board) append_col(index int) bool {
    // If dim = 3:
    //    e : i , j
    //   ---:---,--
    //    0 : 0 , 0
    //    1 : 0 , 1
    //    2 : 0 , 2
    //    3 : 1 , 0
    //    4 : 1 , 1
    //    5 : 1 , 2
    var j int
    j = index / self.dim
    self.cols[j] = append(self.cols[j], index)
    if len(self.cols[j]) == self.dim {
        self.done = true
        self.win_index = self.cols[j]
    }
    return self.done
}

func (self *Board) get_value() int {
    var out int
    for _, v := range self.win_index {

        num, _:= strconv.Atoi(self.row_list[v])
        out += num
    }
    mul, _ := strconv.Atoi(self.last)
    out = out * mul
    return out
}


// Clean board
func (self *Board) clean_raw(board string) []string {
    // join the matrix in one row
    board = strings.ReplaceAll(board, "\n", " ")

    // left only one space as separator
    for strings.Contains(board, "  ") {
        board = strings.ReplaceAll(board, "  ", " ")
    }
    board = strings.Trim(board, " ")

    // split it all elements
    board_list := strings.Split(board, " ")
    return board_list
}

// Initialize with size
func (self *Board) init(size int) {
    self.dim = size
    self.done = false
    self.iters = 0

    var i int = 0
    for i < size {
        i++
        self.cols = make([][]int, size)
        self.rows = make([][]int, size)
    }
}


type Game struct {

    best int
    numbers []string
    board_list []*Board

}


// Build a Bingo Game
func (self *Game) build(file string) {

    for strings.Contains(file, "  ") {
        file = strings.ReplaceAll(file, "  ", " ")
    }
    game_list := strings.Split(file, "\n\n")

    // NUMBERS
    // In the first row there are the randomn numbers.
    self.numbers = strings.Split(strings.Trim(game_list[0], "\n"), ",")

    // BOARDS
    // below are the 5x5 boards
    board_list_str := game_list[1:]

    // create boards and set board list
    for _, board_str := range board_list_str {

        // build a board
        board := Board{}
        board.set_board(board_str, 5)
        self.board_list = append(self.board_list, &board)
    }
}


func (self *Game) check() int {

    board_list := self.board_list
    numbers := self.numbers
    best_iter := len(numbers)
    best_value := 0
    best_index := len(board_list)

    for ib, board := range board_list {
        i := 0
        for !board.done {
            board.mark_number(numbers[i])
            i++
        }
        if v:=board.get_value(); best_value < v {
            best_index = ib
            best_value = v
        }

    }
    fmt.Println("last", best_index, best_value)

    return best_iter
}


func readFile() string {
    // read file from arguments
    input := os.Args[1]
    b, _ := ioutil.ReadFile(input)
    file := string(b)
    return file

}


func main() {

    game := Game{}
    game.build(readFile())
    game.check()

}
