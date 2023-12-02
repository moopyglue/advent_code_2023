package main

import (
	"bufio"
	"fmt"
	"os"
)

var debug = map[string]bool{ "info":false }

func calc(trans map[string]int)(result int) {

	result = 0

	for _, line := range getlines() {

		pinfo( "====================================" )
		pinfo( "start",line )
		last:=len(line)

		left:=-1
		for i:=0 ; i<last ; i++ {
			ll:=line[i:]
			pinfo(ll)
			for m,n := range trans {
				if len(m) > len(ll) {continue}
				pinfo(" ",m,":",n,":",ll[:len(m)]," ")
				if m == ll[:len(m)]  {
					left=n
					break
				}
			}
			pinfo("")
			if left >= 0  { break }
		}

		right:=-1
		for i:=last ; i>0 ; i-- {
			ll:=line[:i]
			pinfo(ll)
			for m,n := range trans {
				if len(m) > len(ll) {continue}
				pinfo(" ",m,":",n,":",ll[len(ll)-len(m):]," ")
				if m == ll[len(ll)-len(m):]  {
					right=n
					break
				}
			}
			pinfo("")
			if right >= 0  { break }
		}
		result = result + (10*left) + right
		pinfo( left,right,result )

	}
	return

}

func part1()(result int) {

	part1_mapping := map[string]int{
		"0":0, "1":1, "2":2, "3":3, "4":4, "5":5, "6":6, "7":7, "8":8, "9":9,
	}
	result=calc(part1_mapping)
	return
}

func part2()(result int) {

	part2_mapping := map[string]int{
		"zero":0, "one":1, "two":2, "three":3, "four":4,
		"five":5, "six":6, "seven":7, "eight":8 , "nine":9,
		"0":0, "1":1, "2":2, "3":3, "4":4, "5":5, "6":6, "7":7, "8":8, "9":9,
	}
	result=calc(part2_mapping)
	return
}

func getlines() (lines []string) {

	args := os.Args[1:]

	if len(args) > 0 {
		file, _ := os.Open(args[0])
		reader := bufio.NewScanner(file)
		for reader.Scan() {
			lines = append(lines, reader.Text())
		}
		file.Close()
	} else {
		pinfo("reading from STDIN")
		reader := bufio.NewScanner(os.Stdin)
		for reader.Scan() {
			lines = append(lines, reader.Text())
		}
	}
	return

}

func pinfo(params ...interface{}) {
	if debug["info"] { fmt.Println(params) }
}


func main() {
	part1res := part1()
	part2res := part2()
	fmt.Println("part 1 =",part1res)
	fmt.Println("part 2 =",part2res)
	os.Exit(0)
}

