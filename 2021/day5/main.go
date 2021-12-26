/*

0,9 -> 5,9  *  y1 = y2
8,0 -> 0,8
9,4 -> 3,4  *  y1 = y2
2,2 -> 2,1  *  x1 = x2
7,0 -> 7,4  *  x1 = x2
6,4 -> 2,0
0,9 -> 2,9  *  y1 = y2
3,4 -> 1,4  *  y1 = y2
0,0 -> 8,8
5,5 -> 8,2


- An entry like 1,1 -> 1,3 covers points 1,1, 1,2, and 1,3.
- An entry like 9,7 -> 7,7 covers points 9,7, 8,7, and 7,7.

For now, only consider horizontal and vertical lines: lines where either x1 =
x2 or y1 = y2.


   0--------9
 | .......1.. 0
 | ..1....1.. |
 | ..1....1.. |
 | .......1.. |
 X .112111211 |
 | .......... |
 | .......... |
 | .......... |
 | .......... |
 | 222111.... 9

   ----Y----

*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)


type Grid struct {
    grid [][]int
}


func (g *Grid) markAll(file string) {

    for _, segment := range strings.Split(file, "\n") {
        (*g).markSegment(segment)
    }

}

func (g *Grid) markSegment(segment string) {

    // convert segment to uint
    var xy []string
    segList := strings.Split(segment, " -> ")
    xy = strings.Split(segList[0], ",")
    xx1, _ := strconv.ParseUint(xy[0], 10, 0)
    yy1, _ := strconv.ParseUint(xy[1], 10, 0)
    xy = strings.Split(segList[1], ",")
    xx2, _ := strconv.ParseUint(xy[0], 10, 0)
    yy2, _ := strconv.ParseUint(xy[1], 10, 0)

    // sorry for this
    x1 := int(xx1)
    x2 := int(xx2)
    y1 := int(yy1)
    y2 := int(yy2)

    // Some tips:
    // - X are rows (axis 0)
    // - Y are columns (axis 1)

    // fmt.Print(segment)

    if (x1 == x2) && (y1 != y2) {
        // Horizontal Line (X axis is fixed and Y's varies)
        var y0, yn int
        // fmt.Print("*")

        if y1 < y2 {
            y0 = y1
            yn = y2
        } else {
            y0 = y2
            yn = y1
        }

        for yi := y0; yi <= yn; yi++ {
            (*g).grid[yi][x1] += 1
        }

    } else if (x1 != x2) && (y1 == y2) {
        // Vertical Line (X's varies and Y axis is fixed)
        var x0, xn int
        // fmt.Print("*")

        if x1 < x2 {
            x0 = x1
            xn = x2
        } else {
            x0 = x2
            xn = x1
        }

        for xi := x0; xi <= xn; xi++ {
            (*g).grid[y1][xi] += 1
        }
    } else if is45(x1, x2, y1, y2) {
        var xd, yd int
        // fmt.Print("***")

        if x1 < x2 {
            xd = 1
        } else {
            xd = -1
        }

        if y1 < y2 {
            yd = 1
        } else {
            yd = -1
        }

        xi := x1
        yi := y1

        for (xi != x2) && (yi != y2) {
            // fmt.Print("|", xi, yi)

            (*g).grid[yi][xi] += 1

            xi += xd
            yi += yd
        }

    }

    // fmt.Print("\n")

}


func is45(x1 int, x2 int, y1 int, y2 int) bool {
    xd := x1 - x2
    yd := y1 - y2
    if xd < 0 {xd *= -1}
    if yd < 0 {yd *= -1}
    return xd == yd
}

func (g *Grid) Count() int {
    var out int
    grid := (*g).grid
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] > 1 {
                out++
            }
        }
    }
    fmt.Println("Count:", out)
    return out
}


func XYmax(file string) (Xmax int, Ymax int) {

    var x, y  int
    var xx, yy uint64
    var xy []string

    file = strings.Trim(file, "\n")
    for _, segment := range strings.Split(file, "\n") {

        segList := strings.Split(segment, " -> ")
        xy = strings.Split(segList[0], ",")
        xx, _ = strconv.ParseUint(xy[0], 10, 0)
        yy, _ = strconv.ParseUint(xy[1], 10, 0)
        x = int(xx)
        y = int(yy)

        if x > Xmax { Xmax = x }
        if y > Ymax { Ymax = y }

        xy = strings.Split(segList[1], ",")
        xx, _ = strconv.ParseUint(xy[0], 10, 0)
        yy, _ = strconv.ParseUint(xy[1], 10, 0)
        x = int(xx)
        y = int(yy)

        if x > Xmax { Xmax = x }
        if y > Ymax { Ymax = y }
    }

    return Xmax + 1, Ymax + 1
}


func buildGrid(Xmax int, Ymax int) (grid [][]int) {
    var x int
    // x is vertical axis
    for x = 0; x < Xmax; x++ {
        // y is horizontal axis
        row := make([]int, Ymax)
        grid = append(grid, row)
    }

    return grid
}


func (g *Grid) init(file string){
    Xmax, Ymax := XYmax(file)
    // fmt.Println(Xmax, Ymax)
    (*g).grid = buildGrid(Xmax, Ymax)
}


func (g *Grid) Print(){

    fmt.Println("")
    fmt.Println("Grid")
    fmt.Println("----")

    grid := (*g).grid

    fmt.Print("  ")
    for h := 0; h < len(grid[0]); h++ {
        fmt.Print(" ", h)
    }
    fmt.Print("\n")
    fmt.Print("  ")
    for h := 0; h < len(grid[0]); h++ {
        fmt.Print(" ", "-")
    }

    fmt.Print("\n")

    for i, row := range grid {
        fmt.Print(i, "|")
        for _, v := range row {
            fmt.Print(" ", v)
        }
        fmt.Print("\n")

    }
}

func readFile() string {
    path := os.Args[1]
    b, _ := ioutil.ReadFile(path)
    file := strings.Trim(string(b), "\n")
    return file
}


func main() {
    file := readFile()
    grid := Grid{}
    grid.init(file)
    // grid.Print()
    grid.markAll(file)
    grid.Print()
    grid.Count()
}
