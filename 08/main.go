// AOC 2023 - day 7 - Camel Cards

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"sort"
	//"strings"
)

var flags = map[string]bool{"info": true}

type node struct {
	left  string
	right string
}

func main() {

	// turn input data to usable puzzle data structure
	var data = getlines()
	path := data[0]
	desert := map[string]node{}
	for n := 1; n < len(data); n++ {
		x := strings.Split(data[n], " ")
		desert[x[0]] = node{left: x[1], right: x[2]}
	}

	// part 1
	pinfo("PART 1")
	pinfo("locations=", "AAA")
	part1res := counter(desert, path, "AAA", "ZZZ")
	pinfo("   counts=", part1res)

	// part 2
	pinfo("PART 2")
	part2res := int64(0)
	// find all nodes with last char as "A"
	locations := []string{}
	counts := []int64{}
	for k, _ := range desert {
		if k[2:] == "A" {
			locations = append(locations, k)
			counts = append(counts, int64(0))
		}
	}
	pinfo("locations=", locations)
	// count the loop size for each location
	for n := 0; n < len(locations); n++ {
		counts[n] = counter(desert, path, locations[n], "Z")
	}
	pinfo("   counts=", counts)
	// calculate the lowest common multiplier
	part2res = LCM(counts)

	// results
	fmt.Println("part 1 =", part1res)
	fmt.Println("part 2 =", part2res)
}

func counter(desert map[string]node, path, start, end string) (result int64) {
	result = 0
	location := start
MainLoop:
	for {
		for n := 0; n < len(path); n++ {
			if path[n] == 'L' {
				location = desert[location].left
			} else {
				location = desert[location].right
			}
			result++
			if location[3-len(end):] == end {
				break MainLoop
			}
		}
	}
	return

}

// find Least Common Multiple (LCM) via GCD
func LCM(i []int64) int64 {
	a := i[0]
	b := i[1]
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	result := i[0] * i[1] / a
	for n := 0; n < len(i[2:]); n++ {
		result = LCM([]int64{result, i[2+n]})
	}
	return result
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
