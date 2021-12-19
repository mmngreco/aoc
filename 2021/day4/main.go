package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

            list[ielem] = "X"
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
    }
    return self.done
}

// Clean board
func (self *Board) clean_raw(board string) []string {
    // drop trailing spaces an new lines
    board = strings.Trim(board, "\n")
    board = strings.Trim(board, " ")

    // join the matrix in one row
    board = strings.ReplaceAll(board, "\n", " ")

    // left only one space as separator
    for strings.Contains(board, "  ") {
        board = strings.ReplaceAll(board, "  ", " ")
    }

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
    for i<size {
        i++
        self.cols = make([][]int, size)
        self.rows = make([][]int, size)
    }
}


type Game struct {

    best int
    numbers []string
    board_list []Board

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
        self.board_list = append(self.board_list, board)
    }
}


func (self *Game) check() int {

    numbers := self.numbers
    board_list := self.board_list
    var best int

    best = len(numbers)

    for ib, board := range board_list {
        i := 0
        for !board.done {
            board.mark_number(numbers[i])
            i++
        }

        if board.done && best > board.iters {
            best = ib
        }

    }

    fmt.Println("end", best, board_list[best].iters)
    return best
}


func readFile() string {
    // read file from arguments
    input := os.Args[1]
    b, _ := ioutil.ReadFile(input)
    file := string(b)
    return file

}


func main() {

    file := readFile()
    game := Game{}
    game.build(file)
    game.check()

}
