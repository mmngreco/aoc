package main


import (
    "os"
    "fmt"
    "io/ioutil"
    "strings"
)


func main() {

    // read file from arguments
    input := os.Args[1]
    b, _ := ioutil.ReadFile(input)
    file := string(b)
    game_list := strings.Split(file, "\n\n")

    // NUMBERS
    // In the first row there are the randomn numbers.
    ans := strings.Split(strings.Trim(game_list[0], "\n"), ",")
    fmt.Println(ans)

    // BOARDS
    // below are the 5x5 boards
    board_list_str := game_list[1]
    // fmt.Println(v)
    for strings.Contains(board_list_str, "  ") {
        board_list_str = strings.ReplaceAll(board_list_str, "  ", " ")
    }
    board_list_str = strings.Trim(board_list_str, "\n")
    board_list := strings.Split(board_list_str, "\n ")

    for _, a := range ans {
        for i, row := range board_list {
            fmt.Println("")
            fmt.Println("board n", i)
            fmt.Println("row", i, "value", row)
            for j, element := range strings.Split(row, " ") {
                fmt.Println("column, element, number:", j, element, a)
            }
        }
    }

    // Process de board
    // Convert them into a matrix ??
    fmt.Println("")
}
