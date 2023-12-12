// AOC 2023 - day 12 -

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var debug = map[string]bool{"info": true}

type row struct {
	known string
	pat   []int
	diff  int
	whole bool
}

var broken_string = " "
var working_string = "O"

func main() {

	// create padded sttrings for ease of use later
	for n := 0; n < 12; n++ {
		broken_string += broken_string
		working_string += working_string
	}

	fives := true
	pinfo("starting")
	// load source data into 'rows' structure
	rows := []row{}
	for _, v := range getlines() {
		a := strings.Split(v, " ")
		a[0] = strings.Replace(a[0], ".", " ", -1)
		a[0] = strings.Replace(a[0], "?", ".", -1)
		a[0] = strings.Replace(a[0], "#", "O", -1)
		if fives {
			a[0] = a[0] + "." + a[0] + "." + a[0] + "." + a[0] + "." + a[0]
			a[1] = a[1] + "," + a[1] + "," + a[1] + "," + a[1] + "," + a[1]
		}
		b := strings.Split(a[1], ",")
		p := []int{}
		s := 0
		for _, k := range b {
			i, _ := strconv.ParseInt(k, 10, 0)
			p = append(p, int(i))
			s = s + int(i)
		}
		s += len(p) - 1
		rows = append(rows, row{known: a[0], pat: p, diff: len(a[0]) - s, whole: true})
	}

	// for each row calculate the cost
	part1res := 0
	for n := 0; n < len(rows); n++ {
		pinfo(n, rows[n])
		count := count_matches(rows[n])
		part1res += count
		pinfo(count, part1res)
	}

	pinfo("part1", part1res)
}

var cache = map[string]bool{}

func count_matches(r row) (count int) {
	count = 0
	mod := 1
	if r.whole {
		mod = 0
		r.whole = false
	}
	if len(r.pat) == 0 {
		re := regexp.MustCompile(r.known)
		if re.MatchString(broken_string[:len(r.known)]) {
			count++
		}
	} else {
		for n := 0; n <= r.diff; n++ {
			s := broken_string[:n+mod] + working_string[:r.pat[0]]
			tt := s + r.known[:len(s)]
			if qq, ok := cache[tt]; ok {
				if qq {
					count += count_matches(row{known: r.known[len(s):], pat: r.pat[1:], diff: r.diff - n})
				}
			} else {
				re := regexp.MustCompile(r.known[:len(s)])
				if re.MatchString(s) {
					cache[tt] = true
					count += count_matches(row{known: r.known[len(s):], pat: r.pat[1:], diff: r.diff - n})
				} else {
					cache[tt] = false
				}
			}
		}
		r.whole = false
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
