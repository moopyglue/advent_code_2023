// AOC 2023 - day 15 -

package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var debug = map[string]bool{"info": true}

type lense struct {
	label string
	focal int
}

func main() {

	log.SetFlags(log.Lmicroseconds)
	pinfo("starting")

	var data = getlines()
	pats := strings.Split(data[0], ",")

	//initialize count for part 1
	part1res := 0

	// initialize boxes for part2
	var p2 = map[int][]lense{}
	for n := 0; n < 256; n++ {
		p2[n] = []lense{}
	}

	for k := range pats {

		// calculate hash of label + whole string + lense id
		operation := pats[k][len(pats[k])-1:]
		label := pats[k][:len(pats[k])-1]
		if operation != "-" {
			label = label[:len(label)-1]
		}
		linecnt := 0
		labelhash := 0
		for n := 0; n < len(pats[k]); n++ {
			if pats[k][n] == '-' || pats[k][n] == '=' {
				labelhash = linecnt
			}
			linecnt += int(pats[k][n])
			linecnt *= 17
			linecnt = linecnt % 256
		}
		// pinfo(operation, label, pats[k], linecnt, labelhash)
		part1res += linecnt

		// place and remove labels as needed
		done := false
		if operation == "-" {
			// remove lense
			for n := 0; n < len(p2[labelhash]); n++ {
				if p2[labelhash][n].label == label {
					p2[labelhash] = append(p2[labelhash][:n], p2[labelhash][n+1:]...)
					break
				}
			}
			continue
		} else {
			//insert/replace lense
			lens, _ := strconv.Atoi(operation)
			for n := 0; n < len(p2[labelhash]); n++ {
				if p2[labelhash][n].label == label {
					tmp := append(p2[labelhash][:n], lense{label: label, focal: lens})
					p2[labelhash] = append(tmp, p2[labelhash][n+1:]...)
					done = true
					break
				}
			}
			if !done {
				p2[labelhash] = append(p2[labelhash], lense{label: label, focal: lens})
			}

		}
	}
	part2res := 0
	for n := 0; n < len(p2); n++ {
		if len(p2[n]) > 0 {
			for q := 0; q < len(p2[n]); q++ {
				part2res += (n + 1) * (q + 1) * p2[n][q].focal
				// pinfo("box", n, "=", p2[n], part2res)
			}
		}
	}

	pinfo("part1", part1res)
	pinfo("part2", part2res)

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
		log.Println(params)
	}
}
