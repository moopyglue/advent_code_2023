// AOC 2023 - day 11 - Cosmic Expansion

package main

import (
	"bufio"
	"log"
	"os"
)

var debug = map[string]bool{"info": true}

func main() {

	pinfo("started")
	//data := getlines()
	var part1res = 0
	var part2res = 0

	pinfo("stopped")
	pinfo("part1", part1res)
	pinfo("part2", part2res)

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
		log.Println(params)
	}
}
