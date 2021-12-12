// day 3
package main

import (
    "fmt"
    "strings"
    "io/ioutil"
)


func readFile(fname string) (nums []string, err error) {
    // read file

    b, err := ioutil.ReadFile(fname)
    if err != nil { return nil, err }

    lines := strings.Split(strings.Trim(string(b), "\n"), "\n")

    return  lines, err
}


func main() {
    file, _ := readFile("sample")
    fmt.Println(file)
}
