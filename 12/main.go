// AOC 2023 - day 12 - Hot Springs

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var debug = map[string]bool{"info": false, "debug": false}

type row struct {
	line string
	cnts []int
}

func main() {

	data := getlines()
	pinfo("starting")
	// load source data into 'rows' structure
	rows := []row{}
	rows5 := []row{}
	for _, v := range data {

		// single copy - part 1
		a := strings.Split(v, " ")
		b := strings.Split(a[1], ",")
		p := []int{}
		s := 0
		for _, k := range b {
			i, _ := strconv.ParseInt(k, 10, 0)
			p = append(p, int(i))
		}
		s += len(p) - 1
		rows = append(rows, row{line: a[0], cnts: p})

		// 5x copy - part 2
		a5 := strings.Split(v, " ")
		a5[0] = a5[0] + "?" + a5[0] + "?" + a5[0] + "?" + a5[0] + "?" + a5[0]
		a5[1] = a5[1] + "," + a5[1] + "," + a5[1] + "," + a5[1] + "," + a5[1]
		b = strings.Split(a5[1], ",")
		p = []int{}
		s = 0
		for _, k := range b {
			i, _ := strconv.ParseInt(k, 10, 0)
			p = append(p, int(i))
		}
		s += len(p) - 1
		rows5 = append(rows5, row{line: a5[0], cnts: p})

	}

	// for each row calculate the cost
	part1res := 0
	for n := 0; n < len(rows); n++ {
		pinfo(rows[n])
		count := count_matches(rows[n].line, rows[n].cnts)
		part1res += count
		pinfo(n+1, count, part1res)
	}
	fmt.Println("part1 =", part1res)

	// for each row calculate the cost
	part2res := 0
	for n := 0; n < len(rows5); n++ {
		pinfo(rows5[n])
		count := count_matches(rows5[n].line, rows5[n].cnts)
		part2res += count
		pinfo(n+1, count, part2res)
	}

	fmt.Println("part2 =", part2res)
}

var cache = map[string]int{}

func count_matches(line string, cnts []int) int {

	xcount := 0
	if line == "" {
		if len(cnts) == 0 {
			xcount = 1
		}
	} else if len(cnts) == 0 {
		if !strings.Contains(line, "#") {
			xcount = 1
		}
	} else {

		key := fmt.Sprint(line, cnts)
		if k, ok := cache[key]; ok {
			return k
		} else {

			k := line[0]
			if k == '?' || k == '.' {
				xcount += count_matches(line[1:], cnts[:])
			}

			if k == '?' || k == '#' {
				if len(line) == cnts[0] && !strings.Contains(line[:cnts[0]], ".") {
					xcount += count_matches(line[cnts[0]:], cnts[1:])
				}
				if len(line) > cnts[0] && !strings.Contains(line[:cnts[0]], ".") && line[cnts[0]:cnts[0]+1] != "#" {
					xcount += count_matches(line[cnts[0]+1:], cnts[1:])
				}
			}
			cache[key] = xcount
		}
	}

	return xcount
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

// debug printing for INFO style lines
func pdebug(params ...interface{}) {
	if debug["debug"] {
		fmt.Println(params)
	}
}
