package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
    Unmarked uint = iota
    Marked   uint = iota
)


type Board struct {
    board [5][5]uint
    marked [5][5]uint
    lastNumber uint
    iterations uint
}


func (b *Board) checkHorizontal() bool {
    for r := 0; r < 5; r++ {
        if (*&b.marked[r][0] == Marked) &&
           (*&b.marked[r][1] == Marked) &&
           (*&b.marked[r][2] == Marked) &&
           (*&b.marked[r][3] == Marked) &&
           (*&b.marked[r][4] == Marked) {
               return true
        }
    }
    return false
}


func (b *Board) checkVertical() bool {
    for c := 0; c < 5; c++ {
        if (*&b.marked[0][c] == Marked) &&
           (*&b.marked[1][c] == Marked) &&
           (*&b.marked[2][c] == Marked) &&
           (*&b.marked[3][c] == Marked) &&
           (*&b.marked[4][c] == Marked) {
               return true
        }
    }
    return false
}


func (b *Board) markNumber(num uint) {
    switch num {
        case *&b.board[0][0]: *&b.marked[0][0] = Marked
        case *&b.board[0][1]: *&b.marked[0][1] = Marked
        case *&b.board[0][2]: *&b.marked[0][2] = Marked
        case *&b.board[0][3]: *&b.marked[0][3] = Marked
        case *&b.board[0][4]: *&b.marked[0][4] = Marked
        case *&b.board[1][0]: *&b.marked[1][0] = Marked
        case *&b.board[1][1]: *&b.marked[1][1] = Marked
        case *&b.board[1][2]: *&b.marked[1][2] = Marked
        case *&b.board[1][3]: *&b.marked[1][3] = Marked
        case *&b.board[1][4]: *&b.marked[1][4] = Marked
        case *&b.board[2][0]: *&b.marked[2][0] = Marked
        case *&b.board[2][1]: *&b.marked[2][1] = Marked
        case *&b.board[2][2]: *&b.marked[2][2] = Marked
        case *&b.board[2][3]: *&b.marked[2][3] = Marked
        case *&b.board[2][4]: *&b.marked[2][4] = Marked
        case *&b.board[3][0]: *&b.marked[3][0] = Marked
        case *&b.board[3][1]: *&b.marked[3][1] = Marked
        case *&b.board[3][2]: *&b.marked[3][2] = Marked
        case *&b.board[3][3]: *&b.marked[3][3] = Marked
        case *&b.board[3][4]: *&b.marked[3][4] = Marked
        case *&b.board[4][0]: *&b.marked[4][0] = Marked
        case *&b.board[4][1]: *&b.marked[4][1] = Marked
        case *&b.board[4][2]: *&b.marked[4][2] = Marked
        case *&b.board[4][3]: *&b.marked[4][3] = Marked
        case *&b.board[4][4]: *&b.marked[4][4] = Marked
    }
}

func (b *Board) Sum() uint {
    var out uint = 0
    if *&b.marked[0][0] == Unmarked { out += *&b.board[0][0] }
    if *&b.marked[0][1] == Unmarked { out += *&b.board[0][1] }
    if *&b.marked[0][2] == Unmarked { out += *&b.board[0][2] }
    if *&b.marked[0][3] == Unmarked { out += *&b.board[0][3] }
    if *&b.marked[0][4] == Unmarked { out += *&b.board[0][4] }
    if *&b.marked[1][0] == Unmarked { out += *&b.board[1][0] }
    if *&b.marked[1][1] == Unmarked { out += *&b.board[1][1] }
    if *&b.marked[1][2] == Unmarked { out += *&b.board[1][2] }
    if *&b.marked[1][3] == Unmarked { out += *&b.board[1][3] }
    if *&b.marked[1][4] == Unmarked { out += *&b.board[1][4] }
    if *&b.marked[2][0] == Unmarked { out += *&b.board[2][0] }
    if *&b.marked[2][1] == Unmarked { out += *&b.board[2][1] }
    if *&b.marked[2][2] == Unmarked { out += *&b.board[2][2] }
    if *&b.marked[2][3] == Unmarked { out += *&b.board[2][3] }
    if *&b.marked[2][4] == Unmarked { out += *&b.board[2][4] }
    if *&b.marked[3][0] == Unmarked { out += *&b.board[3][0] }
    if *&b.marked[3][1] == Unmarked { out += *&b.board[3][1] }
    if *&b.marked[3][2] == Unmarked { out += *&b.board[3][2] }
    if *&b.marked[3][3] == Unmarked { out += *&b.board[3][3] }
    if *&b.marked[3][4] == Unmarked { out += *&b.board[3][4] }
    if *&b.marked[4][0] == Unmarked { out += *&b.board[4][0] }
    if *&b.marked[4][1] == Unmarked { out += *&b.board[4][1] }
    if *&b.marked[4][2] == Unmarked { out += *&b.board[4][2] }
    if *&b.marked[4][3] == Unmarked { out += *&b.board[4][3] }
    if *&b.marked[4][4] == Unmarked { out += *&b.board[4][4] }
    return out
}


func (b *Board) Points() uint {
    return (*b).Sum() * (*b).lastNumber
}


func (b *Board) print() {
    var i uint
    fmt.Println("Board")
    fmt.Println("-----")
    for i = 0; i < 5; i++ {
        fmt.Println(
            "\t",
            (*b).board[i][0],
            "\t",
            (*b).board[i][1],
            "\t",
            (*b).board[i][2],
            "\t",
            (*b).board[i][3],
            "\t",
            (*b).board[i][4],
            "\t",
            "",
            "\t",
            (*b).marked[i][0],
            "\t",
            (*b).marked[i][1],
            "\t",
            (*b).marked[i][2],
            "\t",
            (*b).marked[i][3],
            "\t",
            (*b).marked[i][4],
        )
    }
    fmt.Println("last:", (*b).lastNumber)
    fmt.Println("sum:", (*b).Sum())
}

func (b *Board) done() bool {
    return (*b).checkHorizontal() || (*b).checkVertical()
}


func (b *Board) markAll(numArr []uint) uint {
    var num uint

    for i := 0; i < len(numArr); i++ {

        num = numArr[i]
        (*b).markNumber(num)

        if (*b).done() {
            (*b).lastNumber = num
            (*b).iterations = uint(i)
            return num * (*b).Sum()
        }
    }
    return 0
}


func drawNumbers(file string) []uint {

    // array of numbers to check in each board
    guest_arr := strings.Split(strings.Split(file, "\n\n")[0], ",")
    numbers := make([]uint, len(guest_arr))

    for i, num := range guest_arr {
        v, _ := strconv.Atoi(num)
        numbers[i] = uint(v)
    }

    return numbers
}


func buildBoards(file string) []Board {

    boardRaw := strings.Split(file, "\n\n")[1:]
    board_list := []Board{}

    for _, board := range boardRaw {

        // Convert string to matrix of integers
        boardArr := strings.Split(board, "\n")
        var board_matrix [5][5]uint
        var col int
        for row, row_str := range boardArr {
            col = 0
            for _, e := range strings.Split(row_str, " "){
                if len(e) == 0 { continue }
                v, _ := strconv.Atoi(e)
                board_matrix[row][col] = uint(v)
                col++
            }
        }

        // Create a board
        board_list = append(board_list, Board{board: board_matrix})
    }
    return board_list
}


// Read files from arg
func readFile() string {
    // read file from arguments
    input := os.Args[1]
    b, _ := ioutil.ReadFile(input)
    file := string(b)
    return file
}


func main() {

    file := readFile()
    numbers := drawNumbers(file)
    board_list := buildBoards(file)

    var farIter uint = 0
    var points uint = 0

    for i := 0; i < len(board_list); i++ {

        board_list[i].markAll(numbers)

        if farIter < board_list[i].iterations {
            farIter = board_list[i].iterations
            points = board_list[i].Points()
        }
    }
    fmt.Println(farIter, points)
}
