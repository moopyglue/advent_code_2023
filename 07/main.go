// AOC 2023 - day ? - ?

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

var debug = map[string]bool{"info": false}

type hand struct {
	cards  string
	bet    int64
	sorter string
}

func main() {

	var data = getlines()

	// PART 1
	part1res := int64(0)
	var hands = []hand{}
	for _, s := range data {
		k := strings.Split(s, " ")
		hands = append(hands, hand{cards: k[0], bet: i64(k[1]), sorter: gen_sort_value_part1(k[0])})
	}
	sort.Slice(hands, func(i, j int) bool { return hands[i].sorter < hands[j].sorter })
	for k, hand := range hands {
		//pinfo(hand, (int64(k)+1)*hand.bet)
		part1res += (int64(k) + 1) * hand.bet
	}

	fmt.Println("")

	// PART 2
	part2res := int64(0)
	hands = []hand{}
	for _, s := range data {
		k := strings.Split(s, " ")
		hands = append(hands, hand{cards: k[0], bet: i64(k[1]), sorter: gen_sort_value_part2(k[0])})
	}
	sort.Slice(hands, func(i, j int) bool { return hands[i].sorter < hands[j].sorter })
	for k, hand := range hands {
		pinfo(hand, (int64(k)+1)*hand.bet)
		part2res += (int64(k) + 1) * hand.bet
	}

	fmt.Println("part 1 =", part1res)
	fmt.Println("part 2 =", part2res)
}

func gen_sort_value_part1(s string) (r string) {
	r = ""

	// calculate sprtabl string that represents the order of
	// winning hand status e.g. full, house, 5 of  kind, etc...
	m := map[string]int{}
	for n := 0; n < len(s); n++ {
		m[s[n:n+1]]++
	}
	p := []int{}
	for _, v := range m {
		p = append(p, v)
	}
	sort.Ints(p[:])
	sort.Sort(sort.Reverse(sort.IntSlice(p)))
	p = append(p, 0, 0, 0, 0)
	for _, v := range p {
		r = r + strconv.Itoa(v)
	}
	r = r[:5]

	// change card letters to be values that sort based on value
	// this will let
	// change card letters to be values that sort based on value
	// this will let
	s = strings.Replace(s, "A", "E", -1)
	s = strings.Replace(s, "K", "D", -1)
	s = strings.Replace(s, "Q", "C", -1)
	s = strings.Replace(s, "J", "B", -1)
	s = strings.Replace(s, "T", "A", -1)
	r = r + "-" + s
	return
}

func gen_sort_value_part2(s string) (r string) {
	r = ""

	// generate all possible hands using jokers
	list := "23456789TQK"
	possible := []string{}
	for n := 0; n < len(list); n++ {
		possible = append(possible, strings.Replace(s, "J", list[n:n+1], -1))
	}
	//pinfo(s, possible)
	sortpossible := []string{}
	for _, pos := range possible {
		newpos := ""
		m := map[string]int{}
		for n := 0; n < len(pos); n++ {
			m[pos[n:n+1]]++
		}
		q := []int{}
		for _, v := range m {
			q = append(q, v)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(q)))
		q = append(q, 0, 0, 0, 0)
		for _, v := range q {
			newpos = newpos + strconv.Itoa(v)
		}
		newpos = newpos[:5]
		sortpossible = append(sortpossible, newpos)
	}
	sort.Slice(sortpossible, func(i, j int) bool { return sortpossible[i] < sortpossible[j] })
	r = sortpossible[len(sortpossible)-1]
	//pinfo(r)

	// change card letters to be values that sort based on value
	// this will let
	s = strings.Replace(s, "A", "E", -1)
	s = strings.Replace(s, "K", "D", -1)
	s = strings.Replace(s, "Q", "C", -1)
	s = strings.Replace(s, "J", "1", -1)
	s = strings.Replace(s, "T", "A", -1)

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
