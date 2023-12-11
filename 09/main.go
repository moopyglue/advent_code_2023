// AOC 2023 - day 9 - Mirage Maintenance

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"sort"
	//"strings"
)

var flags = map[string]bool{"info": true}

func main() {

	var data = getlines()
	var part1res = int64(0)
	var part2res = int64(0)

	// singlepass through all data
	for _, line := range data {

		// convert line to big integers
		x := strings.Split(line, " ")
		lols := [][]int64{{}}
		for _, v := range x {
			lols[0] = append(lols[0], i64(v))
		}

		// extrapelate till all zeros
		lolcurr := 0
		for {
			newl := []int64{}
			absx := int64(0)
			for n := 1; n < len(lols[lolcurr]); n++ {
				newl = append(newl, lols[lolcurr][n]-lols[lolcurr][n-1])
				absx += abs(newl[len(newl)-1])
			}
			lolcurr++
			lols = append(lols, newl)
			if absx == 0 {
				break
			}
		}

		// part 1 - add a value to the end of each row to find new value
		lols[lolcurr] = append(lols[lolcurr], int64(0))
		for n := lolcurr - 1; n >= 0; n-- {
			t := len(lols[n]) - 1
			lols[n] = append(lols[n], lols[n][t]+lols[n+1][t])
		}
		part1res += lols[0][len(lols[0])-1]

		// part 2 - add a value to the beginig  of each row
		lols[lolcurr] = append([]int64{0}, lols[lolcurr]...)
		for n := lolcurr - 1; n >= 0; n-- {
			nv := lols[n][0] - lols[n+1][0]
			lols[n] = append([]int64{nv}, lols[n]...)
		}
		part2res += lols[0][0]

	}

	// results
	fmt.Println("part 1 =", part1res)
	fmt.Println("part 2 =", part2res)
}

// ABS function is floating point, this is more efficent than a conversion
func abs(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}

// dirt int64 hack to remove need to error check
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
	if flags["info"] {
		fmt.Println(params)
	}
}
