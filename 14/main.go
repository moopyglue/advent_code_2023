// AOC 2023 - day 11 - Cosmic Expansion

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var debug = map[string]bool{"info": true}

var history = map[string]int{}

func main() {

	var data = getlines()
	var max = 1000000000

	part1 := tilt_north(data)
	pinfo("part1", calc_weight(part1))

	// loop through using tilt cycle
	// but fast forward once a repeating pattern is seen
	pass := 0
	for {
		pass++
		key := strings.Join(data, ",")
		if _, found := history[key]; found {
			// when we see repeat of a previous state we use the gap between
			// last time seen and this time to fast forward to the end at that
			// same repeating point
			pinfo("found repeating cycle of", pass-history[key])
			pass = (max - (max-pass)%(pass-history[key]))
			// wipe the history after we have fast forwarded so
			// fast foprward is not triggered again
			history = map[string]int{}
		}
		history[key] = pass

		// tilt and turn 4 times
		for f := 0; f < 4; f++ {
			data = tilt_north(data)
			data = turn90(data)
		}
		if pass == max {
			break
		}
	}
	pinfo("part2", calc_weight(data))

}

func tilt_north(d []string) []string {
	new := d
	for {
		rolled := false
		for y := 1; y < len(new); y++ {
			for x := 0; x < len(new[0]); x++ {
				if new[y][x] == 'O' && new[y-1][x] == '.' {
					new[y-1] = new[y-1][:x] + "O" + new[y-1][x+1:]
					new[y] = new[y][:x] + "." + new[y][x+1:]
					rolled = true
				}
			}
		}
		if !rolled {
			break
		}
	}
	return new
}

// rotate map clockwise by 90 degrees
func turn90(d []string) []string {
	new := []string{}
	for n := 0; n < len(d[0]); n++ {
		new = append(new, "")
	}
	for y := 0; y < len(d); y++ {
		for x := 0; x < len(d[0]); x++ {
			new[x] = d[y][x:x+1] + new[x]
		}
	}
	return new
}

// calculate weight of provided map
func calc_weight(d []string) (total int) {
	for y := 0; y < len(d); y++ {
		for x := 0; x < len(d[0]); x++ {
			if d[y][x] == 'O' {
				total += len(d) - y
			}
		}
	}
	return
}

// returns input as eitrhegr from standard input or uses first
// command line parameter for filename
func getlines() (lines []string) {

	args := os.Args[1:]
	file, _ := os.Open(args[0])
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		lines = append(lines, reader.Text())
	}
	file.Close()
	return lines
}

// debug printing for INFO style lines
func pinfo(params ...interface{}) {
	if debug["info"] {
		fmt.Println(params)
	}
}
