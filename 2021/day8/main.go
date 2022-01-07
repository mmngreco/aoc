package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)


func isUniqueNum(a string) (int) {
    switch len(a) {
        case 2: return 1
        case 4: return 1
        case 3: return 1
        case 7: return 1
        default: return 0
    }
}

func convertSequences(a string) (num int) {
    var numi int

    a_list := strings.Split(a, " ")

    for _, v := range a_list {
        numi = isUniqueNum(v)
        // fmt.Println(v, numi)
        num += numi
    }

    return num
}


func main() {

    file := os.Args[1]
    b, _ := ioutil.ReadFile(file)
    txt := string(b)
    txt = strings.Trim(txt, "\n")
    lines := strings.Split(txt, "\n")
    counter := 0

    for _, line := range lines {
        left_right := strings.Split(line, " | ")
        // fmt.Println(left_right[0])
        // fmt.Println(left_right[1])
        counter += convertSequences(left_right[1])
    }

    fmt.Println(counter)
}
