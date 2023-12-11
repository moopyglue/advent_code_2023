// AOC 2023 - day 10 - Pipe Maze

package main

import (
	"bufio"
	"fmt"
	"os"
)

var flags = map[string]bool{"info": true}

type vector struct {
	x int
	y int
}

// define an output map as global and the a function
// to update it by flipping a single char
var outchart = []string{}

func set_point(x, y int, s string) {
	outchart[x] = outchart[x][:y] + s + outchart[x][y+1:]
}

func main() {

	var data = getlines()
	list := "F7LJ-|"
	vects := []vector{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	// create a blank output chart and find the "S" start char
	var x, y int
	for a := 0; a < len(data); a++ {
		outchart = append(outchart, "")
		for b := 0; b < len(data[0]); b++ {
			outchart[a] += " "
			if data[a][b:b+1] == "S" {
				x = a
				y = b
			}
		}
	}
	pinfo(x, y)

	// grab directional vectors from "S" that we find
	// a usable path
	var dir vector
	var pf = []vector{}
	for _, v := range vects {
		if x+v.x >= 0 && y+v.y >= 0 {
			p := paths(data[x+v.x][y+v.y : y+v.y+1])
			for _, q := range p {
				if q.x+v.x == 0 && q.y+v.y == 0 {
					pf = append(pf, v)
				}
			}
		}
	}

	// work out what char should be on the map
	// and pick vector to get started
	start_char := ""
	for a := 0; a < len(list); a++ {
		lp := paths(list[a : a+1])
		if (lp[0].x == pf[0].x && lp[0].y == pf[0].y && lp[1].x == pf[1].x && lp[1].y == pf[1].y) ||
			(lp[0].x == pf[1].x && lp[0].y == pf[1].y && lp[1].x == pf[0].x && lp[1].y == pf[0].y) {
			start_char = list[a : a+1]
			break
		}
	}
	dir = pf[0]

	// follow the path round and record those charts on the path
	// on the "outchart" output result
	var pathcount = 0
	for {
		x += dir.x
		y += dir.y
		pathcount++
		val := data[x][y : y+1]
		set_point(x, y, val)
		//pinfo(x, y, val)
		if val == "S" {
			set_point(x, y, start_char)
			break // back at start of loop
		}
		//pinfo("==>  x", x, "y", y, "val", val, "count", pathcount)
		for _, c := range paths(val) {
			if dir.x+c.x != 0 || dir.y+c.y != 0 {
				//pinfo("path", c)
				dir = c // found a path forward
				break
			}
			//pinfo("skipped", c)
		}
	}
	pinfo("pathcount=", pathcount)

	// taking a slice horizontally through the map on each line
	// count the number of lines crossed flipping from inside to outside
	// and visa versa every time a line is crossed
	part2res := 0
	for n := 0; n < len(outchart); n++ {

		inside := false
		direction := "none"
		for m := 0; m < len(outchart[n]); m++ {

			val := outchart[n][m : m+1]
			switch {
			case val == "|":
				inside = !inside
			case val == "L":
				direction = "down"
			case val == "F":
				direction = "up"
			case val == " ":
				if inside {
					set_point(n, m, "O")
					part2res++
				}
			case val == "7":
				if direction == "down" {
					inside = !inside
				}
				direction = "none"
			case val == "J":
				if direction == "up" {
					inside = !inside
				}
				direction = "none"
			default:
				// do nothing
			}
		}
	}

	// print the chart and line number
	for k, v := range outchart {
		fmt.Println(v, k)
	}

	// results
	fmt.Println("part 1 =", pathcount/2.0)
	fmt.Println("part 2 =", part2res)

	// results
}

// identify the vectos in and out of each map char
func paths(s string) []vector {
	switch s {
	case "L":
		return []vector{{-1, 0}, {0, 1}}
	case "-":
		return []vector{{0, -1}, {0, 1}}
	case "|":
		return []vector{{-1, 0}, {1, 0}}
	case "J":
		return []vector{{0, -1}, {-1, 0}}
	case "7":
		return []vector{{1, 0}, {0, -1}}
	case "F":
		return []vector{{1, 0}, {0, 1}}
	default: //"."
		return []vector{}
	}

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
