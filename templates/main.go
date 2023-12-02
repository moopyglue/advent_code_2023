// AOC 2023 - day ?

package main

import (
	"bufio"
	"fmt"
	"os"
)

// debug flags
var debug = map[string]bool{"info": true}

func part1() (result int) {

	result = 0
	return

}

func part2() (result int) {

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

func main() {
	part1res := part1()
	part2res := part2()
	fmt.Println("part 1 =", part1res)
	fmt.Println("part 2 =", part2res)
	os.Exit(0)
}
