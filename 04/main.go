// AOC 2023 - day 4 - ?

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var debug = map[string]bool{"info": false}

type card struct {
	cardnum int
	nums    map[string]bool
	wins    []string
}

func main() {

	var data = map[int]card{}
	for cn, l := range getlines() {
		var c = card{cardnum: cn + 1, nums: make(map[string]bool), wins: make([]string, 0)}
		x := strings.Split(l, "|")
		for _, y := range strings.Split(x[1], " ") {
			c.nums[y] = true
		}
		for _, y := range strings.Split(x[0], " ") {
			if _, ok := c.nums[y]; ok {
				c.wins = append(c.wins, y)
			}
		}
		data[cn+1] = c

	}

	part1res := part1(data)
	part2res := part2(data)
	fmt.Println("part 1 =", part1res)
	fmt.Println("part 2 =", part2res)
}

func part1(data map[int]card) (result int) {

	result = 0
	for _, c := range data {
		score := math.Pow(2, float64(len(c.wins)-1))
		if score >= 1 {
			result += int(score)
			pinfo(score, result)
		}
	}
	return

}

func part2(data map[int]card) (result int) {

	toprocess := []int{}
	for x := 1; x <= len(data); x++ {
		toprocess = append(toprocess, data[x].cardnum)
	}

	processing := 0
	for processing < len(toprocess) {
		var c = data[toprocess[processing]]
		pinfo(c)
		for x := 1; x <= len(c.wins); x++ {
			toprocess = append(toprocess, c.cardnum+x)
		}
		processing++
	}
	result = processing
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
