// AOC 2023 - day 11 - Cosmic Expansion

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var debug = map[string]bool{"info": true}

type node struct {
	cost       int
	retcost    int
	visited    bool
	short_node nodeid
}

type nodeid struct {
	y int
	x int
}

func main() {

	pinfo("started")
	data := getlines()
	grid := load_data(data)
	var part2res = 0
	steps := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	todo := []nodeid{{0, 0}}
	grid[0][0].visited = true
	grid[0][0].retcost = grid[0][0].cost

	for len(todo) > 0 {

		newtodo := []nodeid{}
		fmt.Println("")
		pinfo("===> todo", todo)
		for _, v := range todo {

			xnode := v.x
			ynode := v.y

			pinfo(v)
			for n := 0; n < len(steps); n++ {

				if ynode+steps[n][0] >= 0 && ynode+steps[n][0] < len(grid) && xnode+steps[n][1] >= 0 && xnode+steps[n][1] < len(grid[0]) {

					x := xnode + steps[n][1]
					y := ynode + steps[n][0]
					c := grid[ynode][xnode].retcost

					// if more than 3 in a row for prev path then ignore 4th in same direction
					if xnode-steps[n][1] == grid[ynode][xnode].short_node.x && ynode-steps[n][0] == grid[ynode][xnode].short_node.y {
						p1xnode := xnode - steps[n][1]
						p1ynode := ynode - steps[n][0]
						if p1xnode >= 0 && p1xnode < len(grid[0]) && p1ynode >= 0 && p1ynode < len(grid) {
							if p1xnode-steps[n][1] == grid[p1ynode][p1xnode].short_node.x && p1ynode-steps[n][0] == grid[p1ynode][p1xnode].short_node.y {
								p2xnode := p1xnode - steps[n][1]
								p2ynode := p1ynode - steps[n][0]
								if p2xnode >= 0 && p2xnode < len(grid[0]) && p2ynode >= 0 && p2ynode < len(grid) {
									if p2xnode-steps[n][1] == grid[p2ynode][p2xnode].short_node.x && p2ynode-steps[n][0] == grid[p2ynode][p2xnode].short_node.y {
										pinfo("..............4 in a row avoided", steps[n], nodeid{y, x})
										continue
									}
								}
							}
						}
					}

					if grid[y][x].visited {
						if grid[y][x].cost+c < grid[y][x].retcost {
							grid[y][x].retcost = grid[y][x].cost + c
							grid[y][x].short_node = nodeid{x: xnode, y: ynode}
							pinfo("visited", y, x, "shorter", grid[y][x].retcost)
						} else {
							pinfo("visited", y, x)
						}

					} else {
						grid[y][x].retcost = grid[y][x].cost + c
						grid[y][x].short_node = nodeid{x: xnode, y: ynode}
						grid[y][x].visited = true
						newtodo = append(newtodo, nodeid{x: x, y: y})
						pinfo("new", y, x, grid[y][x].retcost)
					}

				}
			}

		}
		todo = newtodo
	}
	d2 := data
	t := nodeid{len(grid) - 1, len(grid[0]) - 1}
	for !(t.x == 0 && t.y == 0) {
		// pinfo(t)
		d2[t.y] = d2[t.y][:t.x] + "." + d2[t.y][t.x+1:]
		t = grid[t.y][t.x].short_node
	}
	d2[0] = "." + d2[0][1:]
	for n := 0; n < len(d2); n++ {
		pinfo(d2[n])
	}
	for n := 0; n < len(grid); n++ {
		pinfo(grid[n])
	}
	pinfo("stopped")

	var part1res = grid[len(grid)-1][len(grid[0])-1].retcost
	pinfo("part1", part1res)
	pinfo("part2", part2res)

}

// load data into [][] of nodes
func load_data(d []string) (res [][]node) {
	count := 0
	for a := 0; a < len(d); a++ {
		res = append(res, []node{})
		for b := 0; b < len(d[a]); b++ {
			i, _ := strconv.Atoi(d[a][b : b+1])
			res[a] = append(res[a], node{cost: i, retcost: -1})
			count++
		}
	}
	pinfo("created nodes", count)
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
		log.Println(params)
	}
}
