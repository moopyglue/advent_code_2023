// AOC 2023 - day ? - ?

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var debug = map[string]bool{"info": true}

var path = []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity"}

type span struct {
	start int64
	end   int64
}

func main() {

	var spans = []span{}
	var newspans = []span{}
	for _, line := range getlines() {

		// ignore some lines
		if line == "" {
			continue
		}

		tokens := strings.Split(line, " ")
		if tokens[1] == "map:" {
			pinfo("..........................................................", line)
			spans = newspans
			newspans = []span{}
			continue
		}

		// convert seeds to initial spans
		if tokens[0] == "seeds:" {
			// read in seeds
			tokens = tokens[1:]
			for len(tokens) > 0 {
				from, _ := strconv.ParseInt(tokens[0], 10, 0)
				count, _ := strconv.ParseInt(tokens[1], 10, 0)
				to := from + (count - 1)
				s := span{start: from, end: to}
				newspans = append(newspans, s)
				tokens = tokens[2:]
			}
			pinfo("seeds:", newspans)
			continue
		}

		fmt.Println("")
		for _, s := range spans {
			pinfo(s)
		}

		var newstart, newcount, newdest, newend, newdiff int64
		newdest, _ = strconv.ParseInt(tokens[0], 10, 0)
		newstart, _ = strconv.ParseInt(tokens[1], 10, 0)
		newcount, _ = strconv.ParseInt(tokens[2], 10, 0)
		newend = newstart + (newcount - 1)
		newdiff = newdest - newstart
		pinfo("========>", line)
		pinfo("========>", newstart, newend, newdiff)

		for _, orig := range spans {
			fmt.Println("checking", orig, "...")
			var tmpspan span
			if newstart > orig.start && newstart <= orig.end {
				pinfo("uppper insert", orig.start, newstart-1)
				tmpspan = span{orig.start, (newstart - 1)}
				exists := false
				for _, v := range newspans {
					if v.start == tmpspan.start && v.end == tmpspan.end {
						exists = true
					}
				}
				if !exists {
					newspans = append(newspans, tmpspan)
				}
			}
			if newend < orig.end && newend >= orig.start {
				pinfo("lower insert", newend+1, orig.end)
				tmpspan = span{newend + 1, orig.end}
				exists := false
				for _, v := range newspans {
					if v.start == tmpspan.start && v.end == tmpspan.end {
						exists = true
					}
				}
				if !exists {
					newspans = append(newspans, tmpspan)
				}
			}
			if newend < orig.start || newstart > orig.end {
				pinfo("keeping", orig)
				tmpspan = orig
				exists := false
				for _, v := range newspans {
					if v.start == tmpspan.start && v.end == tmpspan.end {
						exists = true
					}
				}
				if !exists {
					newspans = append(newspans, tmpspan)
				}
				continue
			}
			xstart := newstart
			xend := newend
			if xstart < orig.start {
				xstart = orig.start
			}
			if xend > orig.end {
				xend = orig.end
			}
			pinfo("Shurunk", xstart, xend, "adding", newdiff)
			newspans = append(newspans, span{xstart + newdiff, xend + newdiff})

		}

	}

	fmt.Println("")
	for _, s := range spans {
		pinfo(s)
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
	if debug["info"] {
		fmt.Println(params)
	}
}
