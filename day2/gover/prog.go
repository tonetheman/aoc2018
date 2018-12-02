package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func readfile(filename string) []byte {
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("could not read file")
		return nil
	}
	return filedata
}

func stringsToInts(s []string) []int {
	var res []int
	for _, t := range s {
		val, _ := strconv.Atoi(t)
		res = append(res, val)
	}
	return res
}

func savefreq(m map[int]bool) {
	outf, err := os.Create("freq.txt")
	if err != nil {
		fmt.Println("could not open file")
	}
	for k, _ := range m {
		fmt.Fprintf(outf, "%d\n", k)
	}
	outf.Close()
}

func bytesToString(buffer []byte) string {
	return string(buffer)
}

func split(s string) []string {
	if runtime.GOOS == "windows" {
		// might be specific to DOS
		return strings.Split(s, "\r\n")
	} else {
		return strings.Split(s, "\n")
	}
}

func parse2(s string) bool {
	cc := make(map[rune]int)
	for _, char := range s {
		cc[char] += 1
	}
	for _, v := range cc {
		if v == 2 {
			return true
		}
	}
	return false
}

func parse3(s string) bool {
	cc := make(map[rune]int)
	for _, char := range s {
		cc[char] += 1
	}
	for _, v := range cc {
		if v == 3 {
			return true
		}
	}
	return false
}

func count(ss []string) int {
	twocount := 0
	threecount := 0
	for _, s := range ss {
		if parse2(s) {
			twocount++
		}
		if parse3(s) {
			threecount++
		}
	}
	fmt.Println("two", twocount)
	fmt.Println("three", threecount)
	return twocount * threecount
}

func part1() {
	fdata := split(bytesToString(readfile("day2data.txt")))

	//ss := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	res := count(fdata)
	fmt.Println("res", res)

}

func countDiffs(v1, v2 string) int {
	diffcount := 0
	for i, c := range v1 {
		if c == rune(v2[i]) {

		} else {
			diffcount++
		}
	}
	return diffcount
}

func showCommon(v1, v2 string) {
	fmt.Println("showcommond")
	for i, c := range v1 {
		if c == rune(v2[i]) {
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()

}

func part2(fdata []string) {
	index := 0
	for {
		v1 := fdata[index]
		index++
		v2 := fdata[index]
		index++
		//fmt.Println(v1, v2, countDiffs(v1, v2))
		if countDiffs(v1, v2) == 1 {
			fmt.Println("FOUND", v1, v2)
			//showCommon(v1, v2)
		}
		if index > len(fdata)-1 {
			break
		}
	}
}

func part2correct(fdata []string) {
	fmt.Println("starting part2")

	for i := 0; i < len(fdata); i++ {
		for j := 0; j < len(fdata); j++ {
			if i == j {
				//fmt.Println("matched not doing anything...")
				continue
			}
			//fmt.Println("working on", i, j)
			v1 := fdata[i]
			//fmt.Println("after v1")
			v2 := fdata[j]
			//fmt.Println("after v2")
			//fmt.Println(" datum are", v1, v2)
			if countDiffs(v1, v2) == 1 {
				fmt.Println("FOUND", v1, v2)
				showCommon(v1, v2)
			}
		}
	}
}

func main() {
	fdata := split(bytesToString(readfile("day2data.txt")))
	//sort.Strings(fdata)
	//fmt.Println(fdata)
	part2correct(fdata)
}
