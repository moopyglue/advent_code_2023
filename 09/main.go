// AOC 2023 - day ????

package main

import (
	"bufio"
	"fmt"
	"os"
	//"sort"
	//"strings"
)

var flags = map[string]bool{"info": true}

func main() {

	// turn input data to usable puzzle data structure
    // var data = getlines()
    var part1res = int64(0)
    var part2res = int64(0)

	// part 1
	pinfo("PART 1")

	// part 2
	pinfo("PART 2")

	// results
	fmt.Println("part 1 =", part1res)
	fmt.Println("part 2 =", part2res)
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
	if flags["info"] {
		fmt.Println(params)
	}
}
