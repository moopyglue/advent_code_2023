// AOC 2023 - day 12 -

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var debug = map[string]bool{"info": true}

type row struct {
	known string
	pat   []int
	diff  int
}

func main() {

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
		rows = append(rows, row{known: a[0], pat: p, diff: len(a[0]) - s})
	}

	// for each row calculate the cost
	part1res := 0
	for _, r := range rows {
		pinfo(r.known, r.pat, r.minpat, len(r.pat)+1, len(r.known)-r.minpat)
		pinfo(len(options_list(len(r.known)-r.minpat, len(r.pat)+1)))
		// re := regexp.MustCompile(r.known)
		// for _, s := range make_solutions(r) {
		// 	if re.MatchString(s) {
		// 		//pinfo("--->", s)
		// 		part1res++
		// 	}
		// }
	}

	pinfo("part1", part1res)

}

// make all the valid solutions for a specific row
// uses options_list and then converts to string representation
func make_solutions(r row) (result []string) {
	type kk struct {
		c int
		s string
	}
	diff := len(r.known) - r.minpat
	for _, x := range options_list(diff, len(r.pat)+1) {
		tmp := []int{0}
		for _, y := range r.pat {
			tmp = append(tmp, y)
			tmp = append(tmp, 1)
		}
		tmp = append(tmp[:len(tmp)-1], 0)
		for k, v := range x {
			tmp[k*2] = tmp[k*2] + v
		}
		res := ""
		for n := 0; n < (len(tmp) - 2); n = n + 2 {
			for x := 0; x < tmp[n]; x++ {
				res += " "
			}
			for x := 0; x < tmp[n+1]; x++ {
				res += "O"
			}
		}
		for x := 0; x < tmp[len(tmp)-1]; x++ {
			res += " "
		}
		result = append(result, res)
	}
	return
}

// recursive routine to calculate all the perms and combinations
// when dividing a number of balls into a number of cups
func options_list(balls, cups int) (list [][]int) {
	if cups == 1 {
		list = [][]int{{balls}}
	} else {
		for b := 0; b <= balls; b++ {
			k := options_list(balls-b, cups-1)
			for _, x := range k {
				list = append(list, []int{b})
				for _, y := range x {
					list[len(list)-1] = append(list[len(list)-1], y)
				}
			}
		}
	}
	return
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
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
		fmt.Println(params)
	}
}
