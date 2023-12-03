// AOC 2023 - day 3 - Gear Ratios

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var debug = map[string]bool{"info": true}

func main() {

	data := getlines()
	parts := get_parts(data)
	part1res := part1(data, parts)
	part2res := part2(data, parts)
	fmt.Println("part 1 =", part1res)
	fmt.Println("part 2 =", part2res)
}

// struct describes the data for each partnumber
// on the schematic
type part struct {
	x       int
	y       int
	leng    int
	partnum int
}

// scan the schematic and extract all the partnumbers
// to an array
func get_parts(data []string) (part_list []part) {

	last_isnum := false
	partnum := 0
	y_start := 0

	// scan through each charcter of schematic hoovering up part numbers and recording them
	for x := 0; x < len(data); x++ {
		for y := 0; y < len(data[0]); y++ {

			if data[x][y] >= '0' && data[x][y] <= '9' {

				// this is a digit 0-9
				if !last_isnum {
					y_start = y
				}
				n, _ := strconv.Atoi(string(data[x][y]))
				partnum = (partnum * 10) + n
				last_isnum = true

			} else {

				if last_isnum {
					// not a digit and previous char was a digit
					part_list = append(part_list, part{x, y_start, (y - y_start), partnum})
					last_isnum = false
					partnum = 0
				}
			}

		}
	}
	return
}

func part1(data []string, parts []part) (result int) {

	re := regexp.MustCompile("[^0-9.]")

	// for each part found look for surrounding symbols
	for _, n := range parts {
		// build a string containing all chars surrounding the part
		test_str := data[n.x-1][n.y-1:(n.y-1)+n.leng+2] + data[n.x][n.y-1:(n.y-1)+n.leng+2] + data[n.x+1][n.y-1:(n.y-1)+n.leng+2]
		// look for symbols in that string
		if re.Match([]byte(test_str)) {
			result += n.partnum
		}
	}
	return
}

func part2(data []string, parts []part) (result int) {

	// look for all "*" synbols and then match to found part numbers
	for x := 0; x < len(data); x++ {
		for y := 0; y < len(data[0]); y++ {
			if data[x][y] != '*' {
				continue
			}

			// generate a list of parts that touch this star
			tmp_parts := []part{}
			for _, n := range parts {
				if x > n.x+1 || x < n.x-1 {
					continue
				}
				if y > n.y+n.leng || y < n.y-1 {
					continue
				}
				tmp_parts = append(tmp_parts, n)
			}

			// if exactly 2 parts overlapping then mutiey and add to result
			if len(tmp_parts) == 2 {
				result = result + (tmp_parts[0].partnum * tmp_parts[1].partnum)
			}
		}
	}

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
		fmt.Println("INFO:", params)
	}
}
