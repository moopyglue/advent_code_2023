// AOC 2023 - day ? - ?

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

	var data = getlines()
	part1res := part1(data)
	fmt.Println("part 1 =", part1res)
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

func part1(data []string) (result int64) {

	result = 1
	for _, line := range getlines() {
		k := strings.Split(line, " ")
		result = result * win_count(i64(k[0]), i64(k[1]))
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
