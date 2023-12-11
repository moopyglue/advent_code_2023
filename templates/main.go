// AOC 2023 - day 9 - Mirage Maintenance

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var flags = map[string]bool{"info": true}

func main() {

	var data = getlines()
	var part1res = int64(0)
	var part2res = int64(0)

	// results
	fmt.Println("part 1 =", part1res)
	fmt.Println("part 2 =", part2res)
}

// ABS function is floating point, this is more efficent than a conversion
func abs(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}

// dirt int64 hack to remove need to error check
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
	if flags["info"] {
		fmt.Println(params)
	}
}
