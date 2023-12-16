// AOC 2023 - day 11 - Cosmic Expansion

package main

import (
	"bufio"
	"log"
	"os"
)

var debug = map[string]bool{"info": true}

type node struct {
	mirror byte
	down   bool
	right  bool
	up     bool
	left   bool
}

type beam struct {
	direction string
	locx      int
	locy      int
}

func main() {

	pinfo("started")
	data := getlines()

	// PART 1
	part1res := scan4active(data, beam{direction: "right", locx: 0, locy: 0})

	// PART 2
	part2res := 0
	for n := 0; n < len(data[0]); n++ {
		res := scan4active(data, beam{direction: "down", locx: n, locy: 0})
		if res > part2res {
			part2res = res
		}
		res = scan4active(data, beam{direction: "up", locx: n, locy: len(data) - 1})
		if res > part2res {
			part2res = res
		}
	}
	for n := 0; n < len(data); n++ {
		res := scan4active(data, beam{direction: "right", locx: 0, locy: n})
		if res > part2res {
			part2res = res
		}
		res = scan4active(data, beam{direction: "left", locx: len(data[0]) - 1, locy: n})
		if res > part2res {
			part2res = res
		}
	}

	pinfo("stopped")
	pinfo("part1", part1res)
	pinfo("part2", part2res)

}

func scan4active(data []string, bb beam) int {

	var grid = [][]node{}

	// create a grid of data to store the direction which light
	// has left the grid node - where light has already left that node
	// in that direction we stop following the light beam
	for k, v := range data {
		grid = append(grid, []node{})
		for n := 0; n < len(v); n++ {
			grid[k] = append(grid[k], node{mirror: v[n]})
		}
	}

	// array of beams, when no longer any beams we stop
	var beams = []beam{bb}
	for len(beams) > 0 {

		newbeams := []beam{}
		for _, b := range beams {

			// discard beam from "to process list"
			beams = beams[1:]

			// if beam has left grid then ignore it
			if b.locx == -1 || b.locy == -1 || b.locx == len(grid[0]) || b.locy == len(grid) {
				continue
			}

			// caclulate all the new beams generated
			k := node{}
			dir := grid[b.locy][b.locx].mirror
			switch b.direction {
			case "right":
				if dir == '.' || dir == '-' {
					k.right = true
				}
				if dir == '\\' || dir == '|' {
					k.down = true
				}
				if dir == '/' || dir == '|' {
					k.up = true
				}
			case "left":
				if dir == '.' || dir == '-' {
					k.left = true
				}
				if dir == '\\' || dir == '|' {
					k.up = true
				}
				if dir == '/' || dir == '|' {
					k.down = true
				}
			case "down":
				if dir == '.' || dir == '|' {
					k.down = true
				}
				if dir == '\\' || dir == '-' {
					k.right = true
				}
				if dir == '/' || dir == '-' {
					k.left = true
				}
			case "up":
				if dir == '.' || dir == '|' {
					k.up = true
				}
				if dir == '\\' || dir == '-' {
					k.left = true
				}
				if dir == '/' || dir == '-' {
					k.right = true
				}
			}
			if k.left && !grid[b.locy][b.locx].left {
				grid[b.locy][b.locx].left = true
				newbeams = append(newbeams, beam{"left", b.locx - 1, b.locy})
			}
			if k.right && !grid[b.locy][b.locx].right {
				grid[b.locy][b.locx].right = true
				newbeams = append(newbeams, beam{"right", b.locx + 1, b.locy})
			}
			if k.down && !grid[b.locy][b.locx].down {
				grid[b.locy][b.locx].down = true
				newbeams = append(newbeams, beam{"down", b.locx, b.locy + 1})
			}
			if k.up && !grid[b.locy][b.locx].up {
				grid[b.locy][b.locx].up = true
				newbeams = append(newbeams, beam{"up", b.locx, b.locy - 1})
			}
		}
		beams = newbeams
	}

	var count = 0
	for a := 0; a < len(grid); a++ {
		for b := 0; b < len(grid[a]); b++ {
			if grid[a][b].up || grid[a][b].down || grid[a][b].left || grid[a][b].right {
				count++
			}
		}
	}
	return count
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
		log.Println(params)
	}
}
