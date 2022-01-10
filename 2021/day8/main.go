/*

Given
-----

Each digit of a seven-segment display is rendered by turning on or off any of
seven segments named a through g:

  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg


My first idea was convert the string into a binary number in order to avoid
sorting and make easier the mapping between a string (and binary) and integer.

It takes me more than expected to have this running as expected.

The numbers 1, 4, 7 and 8 are the key to dencrypt the others because we can
dencrypt them using their string lenght. So, we need create some rules to
deambiguate the numbers with lenght of 5 and 6.

The first step is taking a string and converting it to a binary as shown in the
following the below table.

Binary matrix
-------------

number  binary            len   string
------  --------------    ---   --------
        1 2 3 4 5 6 7     j
i       a b c d e f g     n

1       0 0 1 0 0 1 0     2     cf       Know it
7       1 0 1 0 0 1 0     3     acf      Know it
4       0 1 1 1 0 1 0     4     bcdf     Know it
3       1 0 1 1 0 1 1     5     acdfg
2       1 0 1 1 1 0 1     5     acdeg
5       1 1 0 1 0 1 1     5     abdfg
6       1 1 0 1 1 1 1     6     abdefg
0       1 1 1 0 1 1 1     6     abcefg
9       1 1 1 1 0 1 1     6     abcdfg
8       1 1 1 1 1 1 1     7     abcdefg  Know it


If we select only numbers with len of 5, we can find the following rules:


number  binary            len   string
------  --------------    ---   --------
2       1 0 1 1 1 0 1     5     acdeg
4       0 1 1 1 0 1 0     4     bcdf
2 | 4 = 1 1 1 1 1 1 1 => all segments active

3       1 0 1 1 0 1 1     5     acdfg
1       0 0 1 0 0 1 0     2     cf
3 | 1 = 1 0 1 1 0 1 1 => equal to 3 (doesn't change)

5       1 1 0 1 0 1 1     5     abdfg


Now, for numbers with len of 6:


number  binary            len   string
------  --------------    ---   --------
6       1 1 0 1 1 1 1     6     abdefg
1       0 0 1 0 0 1 0     2     cf
6 | 1 = 1 1 1 1 1 1 1  => all segments active

0       1 1 1 0 1 1 1     6     abcefg
4       0 1 1 1 0 1 0     4     bcdf
0 | 4 = 1 1 1 1 1 1 1  => all segments active

9       1 1 1 1 0 1 1     6     abcdfg


After having all numbers dencrypted, we have to dencrypt the last 4 numbers.

*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// func debug_group_5() {
//     one := make_digit("ab", 1)
//     eight := make_digit("acedgfb", 8)
//     four := make_digit("eafb", 4)
//     seven := make_digit("dab", 7)
//
//     five := make_digit("cdfbe", 5)
//     two := make_digit("gcdfa", 2)
//     three := make_digit("fbcad", 3)
//
//     debug(two, four) // bingo!
//     // debug(three, four)
//     // debug(five, four)
//
//     debug(three, one)
//     debug(five, one) // bingo!
//
//     fmt.Println("")
//     fmt.Println("")
//     fmt.Println(one)
//     fmt.Println(two)
//     fmt.Println(three)
//     fmt.Println(four)
//     fmt.Println(seven)
//     fmt.Println(eight)
//
// }

type Digit struct {
    Int int
    Bin byte
    Str string
    Len int
}


func debug(left Digit, right Digit) {
    // Print out convenient information for debugging.
    fmt.Println("")
    fmt.Printf("left:\t%07b\t%d\t%s\n", left.Bin, left.Int, left.Str)
    fmt.Printf("right:\t%07b\t%d\t%s\n", right.Bin, right.Int, right.Str)
    fmt.Printf("or:\t%07b\n", left.Bin | right.Bin )
    fmt.Println("------------------------------------------")
}



func len2int(lenght int) (out int) {
    // Convert to numbers based on the length of the string.
    switch lenght {
        case 2: out = 1
        case 3: out = 7
        case 4: out = 4
        case 7: out = 8
        default: out = -1
    }

    return out
}

func str2bin(s string) (out byte) {
    // Convert a string into a binary number.
    // Example:
    // convert("ac")
    // Out: 101
    var base byte = 'a'

    for _, c := range s {
        out += 0b1 << (byte(c)-base)
    }

    return out

}

func Sum(arr []int) (out int) {
    // Sum up the elements of arr.
    // Example:
    // Sum([1,2,3])
    // Out: 6
    for _, v := range arr {
        out += v
    }
    return out
}


func make_digit(s string, num int) Digit {
    // Build a Digit to store relevant information once.
    n := len(s)
    if num == -1 {
        num = len2int(n)
    }

    d := Digit{Bin: str2bin(s), Int: num, Len: n, Str: s}
    return d
}


func dencrypt(s []string, str2int map[byte]int) (out int) {
    // Dencrypt a number using a map.
    var num string = ""
    var bin byte

    for _, v := range s {
        bin = str2bin(v)
        num += strconv.Itoa(str2int[bin])
    }
    out, _ = strconv.Atoi(num)

    return out
}


func make_dencrypter(s []string) (map[byte]int) {
    // Build a map with the binary as key and the number as value appling the
    // rules for the group lenght of 5 and 6.
    var iUnknown, number int
    var d Digit
    var int2digit = make(map[int]Digit, 10)
    var bin2int = make(map[byte]int, 10)
    var digitUnknown [6]Digit

    for _, v := range s {
        d = make_digit(v, -1)
        number = d.Int
        if  number == -1 {
            digitUnknown[iUnknown] = d
            iUnknown++
        } else {
            int2digit[number] = d
            bin2int[d.Bin] = d.Int
        }
    }

    one := int2digit[1]
    four := int2digit[4]
    all_active := 0b1111111  // yes, it's a eight

    for _, v := range digitUnknown {

        if v.Len == 5 {
            if (v.Bin | four.Bin) == byte(all_active) {
                bin2int[v.Bin] = 2
            } else if (v.Bin | one.Bin) == byte(v.Bin) {
                bin2int[v.Bin] = 3
            } else {
                bin2int[v.Bin] = 5
            }
        } else {
            if (v.Bin | one.Bin) == byte(all_active) {
                bin2int[v.Bin] = 6
            } else if (v.Bin | four.Bin) == byte(all_active) {
                bin2int[v.Bin] = 0
            } else {
                bin2int[v.Bin] = 9
            }
        }
    }

    return bin2int
}


func solve_it() {

    file := os.Args[1]
    b, _ := ioutil.ReadFile(file)
    txt := string(b)
    txt = strings.Trim(txt, "\n")
    lines := strings.Split(txt, "\n")

    var out, number int
    var ten_numbers, four_numbers []string
    var dencrypter map[byte]int

    for _, line := range lines {
        left_right := strings.Split(line, " | ")

        ten_numbers = strings.Split(left_right[0], " ")
        dencrypter = make_dencrypter(ten_numbers)
        // fmt.Println(dencrypter)

        four_numbers = strings.Split(left_right[1], " ")
        number = dencrypt(four_numbers, dencrypter)
        // fmt.Printf("%s\t%s\t%d\n", ten_numbers[:], four_numbers[:], numbers)
        out += number
    }

    fmt.Println(out)
}


func main() {
    solve_it()
}
