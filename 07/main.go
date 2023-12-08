// AOC 2023 - day 7 - Camel Cards

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var debug = map[string]bool{"info": true}

type hand struct {
	cards string
	bet   int64
	sort1 string
	sort2 string
}

func main() {

	var data = getlines()
	part1res := int64(0)
	part2res := int64(0)

	// generate a list of hands and store different sort values depending on
	// part1 or part2 rules in use.
	var hands = []hand{}
	for _, s := range data {
		k := strings.Split(s, " ")
		hands = append(hands, hand{cards: k[0], bet: i64(k[1]), sort1: gen_sort_value_part1(k[0]), sort2: gen_sort_value_part2(k[0])})
	}

	// sort using PART 1 rules
	sort.Slice(hands, func(i, j int) bool { return hands[i].sort1 < hands[j].sort1 })
	for k, hand := range hands {
		part1res += (int64(k) + 1) * hand.bet
	}

	// sort using PART 2 rules
	sort.Slice(hands, func(i, j int) bool { return hands[i].sort2 < hands[j].sort2 })
	for k, hand := range hands {
		//pinfo(hand)
		part2res += (int64(k) + 1) * hand.bet
	}

	fmt.Println("part 1 =", part1res)
	fmt.Println("part 2 =", part2res)
}

// generates a sortable hand version usinhg part 1 rules
func gen_sort_value_part1(s string) (r string) {

	// calculate hand type and add softwable version of hand
	r = handtype(s) + s
	r = translate(r, "TJQKA", "ABCDE")
	return

}

// generates a sortable hand version usinhg part 2 rules
func gen_sort_value_part2(s string) (r string) {

	// generate all possible hands using jokers
	// then sort the list of Joker-ized hands and select the best hand (last)
	list := "23456789TQKA"
	possible := []string{}
	for n := 0; n < len(list); n++ {
		possible = append(possible, handtype(strings.Replace(s, "J", list[n:n+1], -1)))
		//pinfo(s, strings.Replace(s, "J", list[n:n+1], -1), handtype(strings.Replace(s, "J", list[n:n+1], -1)))
	}
	sort.Slice(possible, func(i, j int) bool { return possible[i] < possible[j] })
	r = possible[len(possible)-1] + s

	// change card letters to be values that sort based on value
	// this will let
	r = translate(r, "TJQKA", "A1CDE")

	return
}

// basically dirty hack of unix 'tr' command
func translate(s, from, to string) string {
	result := ""
	for n := 0; n < len(s); n++ {
		found := false
		for j := 0; j < len(from); j++ {
			if s[n] == from[j] {
				found = true
				result += to[j : j+1]
			}
		}
		if !found {
			result += s[n : n+1]
		}
	}
	return result
}

// returns the hand type based on hand provided
func handtype(hand string) string {

	m := map[string]int{}
	q := []int{0, 0, 0, 0}
	for n := 0; n < len(hand); n++ {
		m[hand[n:n+1]]++
	}
	for _, v := range m {
		q = append(q, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(q)))
	switch {
	case q[0] == 5:
		return "7"
	case q[0] == 4:
		return "6"
	case q[0] == 3 && q[1] == 2:
		return "5"
	case q[0] == 3 && q[1] == 1:
		return "4"
	case q[0] == 2 && q[1] == 2:
		return "3"
	case q[0] == 2 && q[1] == 1:
		return "2"
	default:
		return "1"
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

// debug printing for INFO style lines
func jinfo(params ...interface{}) {
	s, _ := json.MarshalIndent(params, "", "\t")
	fmt.Println(s)
}
