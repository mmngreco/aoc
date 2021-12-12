package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
    // "reflect"
)


 type Ship struct {
     horizontal int
     vertical int
     depth int
 }


func str2int(movement string) int {
    // split and convert movement string into int.
    splitted := strings.Fields(movement)
    sdir := splitted[0]
    smod := splitted[1]
    mod, _ := strconv.Atoi(smod)

    switch sdir {
    case "forward":
        return mod
    case "down":
        return mod * -1
    case "up":

    }
}

func (s *Ship) move_horizontal(x int) {
    (*s).horizontal += x
}



func main() {
    fmt.Println()
}
