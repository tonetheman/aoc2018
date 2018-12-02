package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	// might be specific to DOS
	return strings.Split(s, "\n")
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

func main() {
	fdata := split(bytesToString(readfile("day2data.txt")))

	//ss := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	res := count(fdata)
	fmt.Println("res", res)
}
