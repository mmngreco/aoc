package main

import (
	"fmt"
	"os"
	"strings"
    "container/list"
)


var (
    // opening
    op1 string = "("
    op2 string = "["
    op3 string = "{"
    op4 string = "<"
    // closing
    cl1 string = ")"
    cl2 string = "]"
    cl3 string = "}"
    cl4 string = ">"
)


type Stack struct {
    stack *list.List
}


func (c *Stack) Push(value string) {
    c.stack.PushFront(value)
}


func (c *Stack) Pop() error {
    if c.stack.Len() > 0 {
        ele := c.stack.Front()
        c.stack.Remove(ele)
    }
    return fmt.Errorf("Pop Error: Stack is empty")
}


func (c *Stack) Front() (string, error) {
    if c.stack.Len() > 0 {
        if val, ok := c.stack.Front().Value.(string); ok {
            return val, nil
        }
        return "", fmt.Errorf("Peep Error: Stack Datatype is incorrect")
    }
    return "", fmt.Errorf("Peep Error: Stack is empty")
}


func (c *Stack) Size() int {
    return c.stack.Len()
}


func (c *Stack) Empty() bool {
    return c.stack.Len() == 0
}


func readFile() (lines []string) {
    fname := os.Args[1]
    b, _ := os.ReadFile(fname)
    txt := strings.Trim(string(b), "\n")
    lines = strings.Split(txt, "\n")
    return lines
}


func is_opening(c string) (bool) {
    switch c {
        case op1: return true
        case op2: return true
        case op3: return true
        case op4: return true
        default: return false
    }
}



func get_closing(o string) string {
    switch o {
        case op1: return cl1
        case op2: return cl2
        case op3: return cl3
        case op4: return cl4
        default: return ""
    }
}


func char2point(c string) int {
    switch c {
        case ")": return 3
        case "]": return 57
        case "}": return 1197
        case ">": return 25137
        default: return 0
    }
}


func parse(line string) int {
    char_list := strings.Split(line, "")
    openings := &Stack{
        stack: list.New(),
    }

    for _, char := range char_list {

        if is_opening(char) {
            openings.Push(char)
            continue
        }

        op, _ := openings.Front()
        cl_expected := get_closing(op)
        if cl_expected == char {
            openings.Pop()
            continue
        }

        return char2point(char)
    }
    return 0
}


func main() {
    lines := readFile()
    total := 0
    for _, line := range lines{
        points := parse(line)
        total += points
    }
    fmt.Println(total)
}
