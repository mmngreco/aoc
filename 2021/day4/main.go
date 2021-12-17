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
    out := strings.Split(file, "\n\n")

    // exlore content
    for i, v := range out {
        if i == 0 {
            // NUMBERS
            ans := strings.Split(v, ",")
            fmt.Println(ans)
        } else {
            // BOARDS
            fmt.Println(v)
            // Process de board
            // Convert them into a matrix ??
        }
        fmt.Println("")
    }
}
