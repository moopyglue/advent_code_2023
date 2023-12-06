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

type relationship struct {
	source int64
	dest   int64
	count  int64
}

type resmap struct {
	from string
	to   string
	list []relationship
}

func main() {

	// read in the data
	// go makes hard work of this as you have to do sepcific conversions
	var almanac = map[string]resmap{}
	var seeds = []int64{}
	var map_from = ""
	var map_to = ""
	for _, line := range getlines() {

		if line == "" {
			continue
		}

		// dealing with header lines of config
		tokens := strings.Split(line, " ")
		if tokens[0] == "seeds:" {
			// read in seeds
			for k, v := range tokens {
				if k != 0 {
					seeds = append(seeds, i64(v))
				}
			}
			pinfo("seeds:", seeds)
			continue
		}
		if tokens[1] == "map:" {
			// create new resmap for new set of relationships
			tok3 := strings.Split(tokens[0], "-")
			map_from = tok3[0]
			map_to = tok3[2]
			continue
		}

		entry, ok := almanac[map_from]
		if !ok {
			entry = resmap{from: map_from, to: map_to, list: []relationship{}}
		}
		entry.list = append(entry.list, relationship{source: i64(tokens[1]), dest: i64(tokens[0]), count: i64(tokens[2])})
		almanac[map_from] = entry

	}
	pinfo("")
	for k, a := range path {
		pinfo(k, almanac[a])
	}

	part1res := part1(almanac, seeds)
	part2res := part2(almanac, seeds)
	fmt.Println("part 1 =", part1res)
	fmt.Println("part 2 =", part2res)
}

func part1(almanac map[string]resmap, seeds []int64) (result int64) {

	result = 0
	for _, s := range seeds {
		for _, p := range path {
			for _, rel := range almanac[p].list {
				if s >= rel.source && s < rel.source+rel.count {
					s += (rel.dest - rel.source)
					break
				}
			}
		}
		pinfo("location=", s)
		if result == 0 || s < result {
			result = s
		}

	}
	return

}

func part2(almanac map[string]resmap, seeds []int64) (result int64) {

	result = 0

	for len(seeds) > 0 {

		from := seeds[0]
		cnt := seeds[1]
		pinfo("=====>", from, cnt)
		for s := from; s < (from + cnt); s++ {
			calc := s
			for _, p := range path {
				for _, rel := range almanac[p].list {
					if calc >= rel.source && calc < rel.source+rel.count {
						calc += (rel.dest - rel.source)
						break
					}
				}
			}
			if result == 0 || calc < result {
				result = calc
			}
		}
		seeds = seeds[2:]

	}
	return

}

func i64(s string) (i int64) {
	i, _ = strconv.ParseInt(s, 10, 0)
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
