// AOC 2023 - day ? - ?

package main

import (
	"bufio"
	"fmt"
	"os"
)
var debug = map[string]bool{"info": true}

func main() {

	var data = getlines()
	part1res := part1(data)
	part2res := part2(data)
	fmt.Println("part 1 =", part1res)
	fmt.Println("part 2 =", part2res)
}

func part1(data []string) (result int) {

	result = 0
	for _, line := range getlines() {
		pinfo(line)
	}
	return

}

func part2(data []string) (result int) {

	result = 0
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

