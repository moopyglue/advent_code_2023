package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// debug flags
var debug = map[string]bool{"info": false}

// part1
func part1() (result int) {

	allcounts := make(map[string]int)
	result = 0

	re := regexp.MustCompile(" *[:,;] *")
	for _, line := range getlines() {

		// extrat data from line
		split := re.Split(line, -1)
		if len(split) < 2 {
			continue
		}

		// process header line
		if split[0] == "all" {
			pinfo("header found  ")
			for _, v := range split[1:] {
				k := strings.Split(v, " ")
				allcounts[k[1]], _ = strconv.Atoi(k[0])
			}
			bytearray, _ := json.Marshal(allcounts)
			pinfo(string(bytearray))
			continue
		}

		// process data line
		failed := false
		for _, v := range split[1:] {
			k := strings.Split(v, " ")
			if i, _ := strconv.Atoi(k[0]); allcounts[k[1]] < i {
				pinfo("maxed", i, v)
				failed = true
				break
			}
		}
		if failed {
			pinfo("failed")
			continue
		}
		k := strings.Split(split[0], " ")
		i, _ := strconv.Atoi(k[1])
		result = result + i
		pinfo(split[0], result)

	}
	pinfo("part 1 result =", result)
	return

}

// part 2
func part2() (result int) {

	pinfo("part2")
	result = 0

	re := regexp.MustCompile(" *[:,;] *")
	for _, line := range getlines() {

		// extrat data from line
		split := re.Split(line, -1)
		if len(split) < 2 {
			continue
		}
		if split[0] == "all" {
			continue
		}
		pinfo(line)

		// process data line
		maxs := make(map[string]int)
		for _, v := range split[1:] {
			k := strings.Split(v, " ")
			cube_color := k[1]
			cube_count, _ := strconv.Atoi(k[0])
			if _, ok := maxs[cube_color]; !ok {
				maxs[cube_color] = cube_count
			} else {
				if maxs[cube_color] < cube_count {
					maxs[cube_color] = cube_count
				}
			}
		}
		bytearray, _ := json.Marshal(maxs)
		pinfo(string(bytearray))
		force := 1
		for _, v := range maxs {
			force = force * v
		}
		pinfo("force=", force)
		result = result + force
	}

	pinfo("part 2 result =", result)
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

func main() {
	part1res := part1()
	part2res := part2()
	fmt.Println("part 1 =", part1res)
	fmt.Println("part 2 =", part2res)
	os.Exit(0)
}
