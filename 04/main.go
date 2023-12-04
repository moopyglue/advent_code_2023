// AOC 2023 - day 4 - Scratchcards

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
	score   int
}

func main() {

	var data = map[int]card{}

	// build a map of all the cards
	// NOTE: source file altered for simpler processing ("Card #:" and superflous spaces removed)
	for cn, l := range getlines() {

		// create a 'card' from data row
		var c = card{cardnum: cn + 1, nums: make(map[string]bool), wins: make([]string, 0), score: 0}
		x := strings.Split(l, "|")

		for _, y := range strings.Split(x[1], " ") {
			c.nums[y] = true // creating lookup list
		}

		for _, y := range strings.Split(x[0], " ") {
			if _, ok := c.nums[y]; ok {
				c.wins = append(c.wins, y) // creating win list
			}
		}

		// count the matching wins
		c.score = len(c.wins)
		data[cn+1] = c
	}

	part1res := part1(data)
	part2res := part2(data)
	fmt.Println("part 1 =", part1res)
	fmt.Println("part 2 =", part2res)
}

func part1(data map[int]card) (result int) {

	// review each card and calculate a total score
	result = 0
	for _, c := range data {
		score := math.Pow(2, float64(c.score-1))
		if score >= 1 {
			result += int(score)
		}
	}
	return

}

func part2(data map[int]card) (result int) {

	// Create a list of card numbers that need processed
	// to_process[] is a list of card numbers rather than actual cards
	// as there will be a LOT of cards
	to_process := []int{}
	for x := 1; x <= len(data); x++ {
		to_process = append(to_process, data[x].cardnum)
	}

	// process each card, adding cards to the end of the list
	// as we continue.
	processing := 0
	for processing < len(to_process) {
		// var c = data[to_process[processing]]
		// pinfo(data[to_process[processing]])
		for x := 1; x <= data[to_process[processing]].score; x++ {
			to_process = append(to_process, data[to_process[processing]].cardnum+x)
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
