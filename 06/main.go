// AOC 2023 - day 6 - Wait For It

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var debug = map[string]bool{"info": true}

func main() {

	data := getlines()

	part2res := int64(1)
	for _, line := range data {
		k := strings.Split(line, " ")
		part2res = part2res * win_count(i64(k[0]), i64(k[1]))
	}
	fmt.Println("part 2 =", part2res)
}

func win_count(length, record int64) (result int64) {
	result = 0
	for n := int64(1); n < length; n++ {
		dist := n * (length - n)
		if dist > record {
			result++
		}
	}
	return
}

// quick and dirty string to int64 conversion to make code more
// readable
func i64(s string) (i int64) {
	i, _ = strconv.ParseInt(s, 10, 0)
	return
}

// returns input as eitrhegr from standard input or uses first
// command line parameter for filename
func getlines() (lines []string) {

	args := os.Args[1:]

	if len(args) > 0 {
		// use filename provided
		file, _ := os.Open(args[0])
		reader := bufio.NewScanner(file)
		for reader.Scan() {
			lines = append(lines, reader.Text())
		}
		file.Close()
	} else {
		// use STDIN
		pinfo("reading from STDIN")
		reader := bufio.NewScanner(os.Stdin)
		for reader.Scan() {
			lines = append(lines, reader.Text())
		}
	}
	return

}

// debug printing for INFO style lines
func pinfo(params ...interface{}) {
	if debug["info"] {
		fmt.Println(params)
	}
}
