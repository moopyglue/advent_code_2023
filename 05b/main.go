// AOC 2023 - day ? - ?
// ──────────────────────────────────────
// ──────────╔╗──╔═╗───────╔╗──────╔╗────
// ─────────╔╝╚╗─║╔╝───────║║──────║║────
// ────╔═╗╔═╩╗╔╝╔╝╚╦╦═╗╔╦══╣╚═╦══╦═╝║────
// ────║╔╗╣╔╗║║─╚╗╔╬╣╔╗╬╣══╣╔╗║║═╣╔╗║────
// ────║║║║╚╝║╚╗─║║║║║║║╠══║║║║║═╣╚╝║────
// ────╚╝╚╩══╩═╝─╚╝╚╩╝╚╩╩══╩╝╚╩══╩══╝────
// ──────────────────────────────────────

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
	diff  int64
}

type resmap struct {
	from string
	to   string
	list []span
}

func main() {

	var almanac = getalmanac()
	pinfo("===== ALMANAC =====")
	for _, s := range path {
		pinfo(almanac[s])
		fmt.Println("")
	}

	// part 1
	master_map := resmap{from: "", to: "", list: []span{{start: 79, end: 93, diff: 0},{start: 55, end: 68, diff: 0}	} }
	pinfo(master_map)
	for _, s := range path {
		master_map = merge_map(master_map, almanac[s])
		fmt.Println("-------------------")
		pinfo(almanac[s])
		pinfo(master_map)
		fmt.Println("")
	}

}

func merge_map(orig, newlayer resmap) (res resmap) {

	var spans = []span{orig}
	var unchanged = []span{}
	var newspans = []span{}

	for _, orig := range spans {

		if newstart > orig.start && newstart <= orig.end {
			pinfo("upper untouched mapped span", orig.start, newstart-1)
			unchanged = append(unchanged, span{orig.start, (newstart - 1),origin.diff})
		}
		if newend < orig.end && newend >= orig.start {
			pinfo("lower untouched mapped span", newend+1, orig.end)
			unchanged = append(unchanged, span{newend + 1, orig.end,orig.diff})
		}
		if newend < orig.start || newstart > orig.end {
			pinfo("keeping original", orig)
			unchanged = unchanged(newspans, orig)
		}

		// a cropped mapped span is stored seperatly until the next
		// map is started or the file runs out
		xstart := newstart
		xend := newend
		if xstart < orig.start {
			xstart = orig.start
		}
		if xend > orig.end {
			xend = orig.end
		}
		pinfo("a cropped mapped span,", xstart, xend, "adding", newdiff)
		tmpspans = append(tmpspans, span{xstart + newdiff, xend + newdiff})

		// saved spans in newspans repalces originals from spans

	return
}

func getalmanac() (almanac map[string]resmap) {

	// read in the data
	// go makes hard work of this as you have to do sepcific conversions
	almanac = map[string]resmap{}
	var map_from = ""
	var map_to = ""
	for _, line := range getlines() {

		tokens := strings.Split(line, " ")
		if line == "" || tokens[0] == "seeds:" {
			// ignore blank lines and seed line
			continue
		} else if tokens[1] == "map:" {
			// set tags for new resmap for new set of relationships
			tok3 := strings.Split(tokens[0], "-")
			map_from = tok3[0]
			map_to = tok3[2]
		} else {
			entry, ok := almanac[map_from]
			if !ok {
				// for first new span of a map create resmap to contain it
				entry = resmap{from: map_from, to: map_to, list: []span{}}
			}
			// add span to existing resmap entry
			entry.list = append(entry.list, span{start: i64(tokens[1]), end: i64(tokens[1]) + (i64(tokens[2]) - 1), diff: i64(tokens[0]) - i64(tokens[1])})
			almanac[map_from] = entry
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
