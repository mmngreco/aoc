package main

import (
	"fmt"
	"os"
	"strings"
)


var (
    op1 string = "()"
    op2 string = "[]"
    op3 string = "{}"
    op4 string = "<>"
)


func has_pair(line string) bool {
    switch {
        case strings.Contains(line, op1): return true
        case strings.Contains(line, op2): return true
        case strings.Contains(line, op3): return true
        case strings.Contains(line, op4): return true
        default: return false
    }
}

func readFile() (lines []string) {
    fname := os.Args[1]
    b, _ := os.ReadFile(fname)
    txt := strings.Trim(string(b), "\n")
    lines = strings.Split(txt, "\n")
    return lines
}

func parse(line string) () {
    fmt.Println("Start\n", line)
    for has_pair(line) {
        line = strings.ReplaceAll(line, op1, "")
        line = strings.ReplaceAll(line, op2, "")
        line = strings.ReplaceAll(line, op3, "")
        line = strings.ReplaceAll(line, op4, "")
        fmt.Println(line)
    }
}

func main() {
    lines := readFile()
    for _, line := range lines{
        parse(line)
    }
    fmt.Println()
}
