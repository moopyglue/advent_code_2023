// AOC 2023 - day 13 - Point of Incidence

package main

import (
	"bufio"
	"fmt"
	"os"
)

var debug = map[string]bool{"info": true}

type valley struct {
	lines []string
}

func main() {

	// read in input data into descrete data
	// valleys list
	var data = getlines()
	var t = []string{}
	var vals = []valley{}
	for _, v := range data {
		if v == "" {
			vals = append(vals, valley{lines: t[:]})
			t = []string{}
		} else {
			t = append(t, v)
		}
	}
	if len(t) > 0 {
		vals = append(vals, valley{lines: t[:]})
	}

	// for each valley calculate the total
	part1res := 0
	part2res := 0
	for _, vnorm := range vals {

		vflip := hoz2vert(vnorm)

		// PART 1
		c := 0
		cnorm := foldline(vnorm)
		for k := range cnorm {
			c += cnorm[k] * 100
		}
		cflip := foldline(vflip)
		for k := range cflip {
			c += cflip[k]
		}
		part1res += c

		// PART 2
		part2res += findsmudge(vnorm)

	}
	pinfo("part1", part1res)
	pinfo("part2", part2res)

}

// place a smudge in each location and then
// look for additional reflections not seen before
func findsmudge(v valley) int {

	// find original fold lines
	bnorm := foldline(v)
	bflip := foldline(hoz2vert(v))

	for a := 0; a < len(v.lines); a++ {
		orig := v.lines[a]
		for b := 0; b < len(v.lines[0]); b++ {

			// for each potential smudge
			if orig[b] == '#' {
				v.lines[a] = orig[0:b] + "." + orig[b+1:]
			} else {
				v.lines[a] = orig[0:b] + "#" + orig[b+1:]
			}
			cnorm := foldline(v)
			cflip := foldline(hoz2vert(v))
			if (len(cnorm) != 0 || len(cflip) != 0) &&
				(!array_cmp(bnorm, cnorm) || !array_cmp(bflip, cflip)) {

				// a smudge has generatedd a new refelction
				if array_cmp(bnorm, cnorm) {
					return find_new(bflip, cflip)
				} else {
					return find_new(bnorm, cnorm) * 100
				}
			}
			// remve tested smudge for next loop
			v.lines[a] = orig

		}
	}
	return 0
}

// finds first new int in b[] that is not in a[]
func find_new(a []int, b []int) int {
	for x := 0; x < len(b); x++ {
		found := false
		for y := 0; y < len(a); y++ {
			if b[x] == a[y] {
				found = true
			}
		}
		if !found {
			return b[x]
		}
	}
	return 0
}

// compares 2 []int arrays, if the same the 'true' else 'false'
func array_cmp(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for n := 0; n < len(a); n++ {
		if a[n] != b[n] {
			return false
		}
	}
	return true
}

// flips a valley diagnoally so that x,y now = y,x
// this enables us to use the foldline() routine without
// alteration for both vertical and horizontal checks
func hoz2vert(v valley) valley {
	x := []string{}
	for n := 0; n < len(v.lines[0]); n++ {
		x = append(x, "")
	}
	for n := 0; n < len(v.lines); n++ {
		for m := 0; m < len(v.lines[0]); m++ {
			x[m] += v.lines[n][m : m+1]
		}
	}
	return valley{lines: x}
}

// scans each potential fold line to see if it is the mirror point
// returns 0 if no mirror point found
func foldline(v valley) []int {

	result := []int{}
	for n := 0; n < len(v.lines)-1; n++ {
		found := true
		for m := 0; (n-m) >= 0 && (n+m)+1 < len(v.lines); m++ {
			if v.lines[n-m] != v.lines[n+1+m] {
				found = false
				break
			}
		}
		if found {
			result = append(result, n+1)
		}
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
	if debug["info"] {
		fmt.Println(params)
	}
}
