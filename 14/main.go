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

	single_tilt := tilt_north(data)
	pinfo("part1", calc_weight(single_tilt))

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
			pass = (max - (max-pass)%(pass-history[key]))
			// wipe the history after we have fast forwarded so
			// fast foprward is not triggered again
			history = map[string]int{}
		}
		data = tilt_east(tilt_south(tilt_west(tilt_north(data))))
		history[key] = pass
		if pass == max {
			break
		}
	}
	pinfo("part2", calc_weight(data))

}

func tilt_north(d []string) []string {
	for {
		rolled := false
		for y := 1; y < len(d); y++ {
			for x := 0; x < len(d[0]); x++ {
				if d[y][x] == 'O' && d[y-1][x] == '.' {
					d[y-1] = d[y-1][:x] + "O" + d[y-1][x+1:]
					d[y] = d[y][:x] + "." + d[y][x+1:]
					rolled = true
				}
			}
		}
		if !rolled {
			break
		}
	}
	return d
}

func tilt_south(d []string) []string {
	for {
		rolled := false
		for y := len(d) - 2; y >= 0; y-- {
			for x := 0; x < len(d[0]); x++ {
				if d[y][x] == 'O' && d[y+1][x] == '.' {
					d[y+1] = d[y+1][:x] + "O" + d[y+1][x+1:]
					d[y] = d[y][:x] + "." + d[y][x+1:]
					rolled = true
				}
			}
		}
		if !rolled {
			break
		}
	}
	return d
}

func tilt_west(d []string) []string {
	for {
		rolled := false
		for y := 0; y < len(d); y++ {
			for x := 1; x < len(d[0]); x++ {
				if d[y][x] == 'O' && d[y][x-1] == '.' {
					d[y] = d[y][:x-1] + "O." + d[y][x+1:]
					rolled = true
				}
			}
		}
		if !rolled {
			break
		}
	}
	return d
}

func tilt_east(d []string) []string {
	for {
		rolled := false
		for y := 0; y < len(d); y++ {
			for x := len(d[0]) - 2; x >= 0; x-- {
				if d[y][x] == 'O' && d[y][x+1] == '.' {
					d[y] = d[y][:x] + ".O" + d[y][x+2:]
					rolled = true
				}
			}
		}
		if !rolled {
			break
		}
	}
	return d
}

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
