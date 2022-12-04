/*

Part 1
------

     2 >1< 9  9  9  4  3  2  1 >0
        ^                       ^
     3  9  8  7  8  9  4  9  2  1
           v
     9  8 >5< 6  7  8  9  8  9  2
           ^
     8  7  6  7  8  9  6  7  8  9
                       v
     9  8  9  9  9  6 >5< 6  7  8

Each number corresponds to the height of a particular location, where 9 is the
highest and 0 is the lowest a location can be.

Your first goal is to find the low points - the locations that are lower than
any of its adjacent locations. Most locations have four adjacent locations (up,
down, left, and right); locations on the edge or corner of the map have three
or two adjacent locations, respectively. (Diagonal locations do not count as
adjacent.)

In the above example, there are four low points, all highlighted: two are in
the first row (a 1 and a 0), one is in the third row (a 5), and one is in the
bottom row (also a 5). All other locations on the heightmap have some lower
adjacent location, and so are not low points.

The risk level of a low point is 1 plus its height. In the above example, the
risk levels of the low points are 2, 1, 6, and 6. The sum of the risk levels of
all low points in the heightmap is therefore 15.


Part 2
------

Next, you need to find the largest basins so you know what areas are most
important to avoid.

A basin is all locations that eventually flow downward to a single low point.
Therefore, every low point has a basin, although some basins are very small.
Locations of height 9 do not count as being in any basin, and all other
locations will always be part of exactly one basin.

The size of a basin is the number of locations within the basin, including the
low point. The example above has four basins.

The top-left basin, size 3:

2199943210
3987894921
9856789892
8767896789
9899965678
The top-right basin, size 9:

2199943210
3987894921
9856789892
8767896789
9899965678
The middle basin, size 14:

2199943210
3987894921
9856789892
8767896789
9899965678
The bottom-right basin, size 9:

2199943210
3987894921
9856789892
8767896789
9899965678

Find the three largest basins and multiply their sizes together. In the above
example, this is 9 * 14 * 9 = 1134.


*/
package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const visited = 9

var (
	//    0 1 2
	//
	// 0  1 2 3
	// 1  1 # 3
	// 2  1 2 3
	origin = [2]int{0, 0}
	up     = [2]int{-1, 0}
	down   = [2]int{1, 0}
	left   = [2]int{0, -1}
	right  = [2]int{0, 1}
)

func is_inside(data [][]int, row_i int, col_i int) bool {
	nrows := len(data)
	ncols := len(data[0])
	return (row_i >= 0) && (row_i < nrows) && (col_i >= 0) && (col_i < ncols)
}

func is_basin(data [][]int, row_i int, col_i int) bool {
	return is_inside(data, row_i, col_i) && (data[row_i][col_i] < visited)
}

func basinFinder(data [][]int, row_i int, col_i int, motion [2]int) (out int) {

	row_i += motion[0]
	col_i += motion[1]

	if !is_basin(data, row_i, col_i) {
		return 0
	}

	data[row_i][col_i] = visited

	out = 1
	out += basinFinder(data, row_i, col_i, up)
	out += basinFinder(data, row_i, col_i, down)
	out += basinFinder(data, row_i, col_i, left)
	out += basinFinder(data, row_i, col_i, right)

	return out
}

func printMatrix(data [][]int) {
	for _, r := range data {
		fmt.Println(r)
	}
}

func make_matrix(nrows int, ncols int) [][]int {
	matrix := make([][]int, nrows)
	for i := range matrix {
		matrix[i] = make([]int, ncols)
	}
	return matrix
}

func lines2matrix(lines []string) [][]int {

	nrows := len(lines)
	ncols := len(lines[0])

	// Create an empty matrix
	matrix := make_matrix(nrows, ncols)

	// convert to int
	for row_i, row_v := range lines {
		for col_i, col_v := range row_v {
			v, _ := strconv.Atoi(string(col_v))
			matrix[row_i][col_i] = int(v)
		}
	}
	return matrix
}

func quicksort_rev(arr []int) (out []int) {
	n := len(arr)

	if n < 2 {
		// early stop
		return arr
	}

	idx := rand.Intn(n)
	pivot := arr[idx]

	var (
		left  []int
		right []int
		equal []int
	)

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		}
		if v > pivot {
			right = append(right, v)
		}
		if v == pivot {
			equal = append(equal, v)
		}
	}

	out = append(out, quicksort_rev(right)...)
	out = append(out, equal...)
	out = append(out, quicksort_rev(left)...)

	return out

}

func readFile() (matrix [][]int) {
	// read file and return a Matrix
	filename := os.Args[1]
	b, _ := ioutil.ReadFile(filename)
	str := strings.Trim(string(b), "\n")
	lines := strings.Split(str, "\n")
	matrix = lines2matrix(lines)
	return matrix
}

func main() {
	matrix := readFile()

	var basinList []int

	// find low points
	for row_i, row_v := range matrix {
		for col_i := range row_v {
			counter := basinFinder(matrix, row_i, col_i, origin)
			if counter > 0 {
				basinList = append(basinList, counter)
			}
		}
	}

	// solution
	sorted := quicksort_rev(basinList)
	fmt.Println("3-largest basin:", sorted[0:3])
	fmt.Println("Solution:", sorted[0]*sorted[1]*sorted[2])

}
