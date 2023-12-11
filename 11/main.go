// AOC 2023 - day 11 - Cosmic Expansion

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var debug = map[string]bool{"info": true}

type node struct {
	x int
	y int
}

func main() {

	var data = getlines()

	// create a list of nodes from the '#'s in the sky map
	foundpoints := []node{}
	var seenx = [1000]bool{}
	var seeny = [1000]bool{}
	for y, yv := range data {
		for x := 0; x < len(yv); x++ {
			if yv[x] == '#' {
				foundpoints = append(foundpoints, node{x, y})
				seenx[x] = true
				seeny[y] = true
			}
		}

	}

	multiplier := 1000000
	// create x and y modifers to double blank lines for time distotion
	modx := [1000]int{}
	for x := 0; x < len(data[0]); x++ {
		if !seenx[x] {
			for t := x; t < len(data[0]); t++ {
				modx[t] += multiplier - 1
			}
		}
	}
	mody := [1000]int{}
	for y := 0; y < len(data); y++ {
		if !seeny[y] {
			for t := y; t < len(data); t++ {
				mody[t] += multiplier - 1
			}
		}
	}

	// add modifers to data points
	points := []node{}
	for _, p := range foundpoints {
		points = append(points, node{p.x + modx[p.x], p.y + mody[p.y]})
	}

	//part1 - calculate distances between each of all the galaxies
	part1res := 0
	for a := 0; a < len(points); a++ {
		for b := a + 1; b < len(points); b++ {
			part1res += abs(points[a].x - points[b].x)
			part1res += abs(points[a].y - points[b].y)
			//pinfo(a, b, points[a], points[b], part1res)
		}
	}

	// pinfo(foundpoints)
	// pinfo(points)
	// pinfo(modx[:len(data[0])])
	// pinfo(mody[:len(data)])

	fmt.Println("part 1 =", part1res)
	//fmt.Println("part 2 =", part2res)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
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

// quick and dirty single value response
// string -> int64 conversion
func i64(s string) (i int64) {
	i, _ = strconv.ParseInt(s, 10, 0)
	return
}

// debug printing for INFO style lines
func pinfo(params ...interface{}) {
	if debug["info"] {
		fmt.Println(params)
	}
}
