package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type instr struct {
	pre, post string
	avail     bool
}

func readfile(filename string) []byte {
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("could not read file")
		return nil
	}
	return filedata
}

func movePre(ipre string, dest *[]instr, src *[]instr) {
	fmt.Println("need to move any pre eq to", ipre)
	for i := range *src {
		if (*src)[i].pre == ipre && (*src)[i].avail {
			*dest = append(*dest, (*src)[i])
			(*src)[i].avail = false
		}
	}
}

func markNotAvail(i string, src *[]instr) {
	fmt.Println("marking", i, "not avail")
	for k := range *src {
		_tmp := (*src)[k]
		if _tmp.pre == i {
			_tmp.avail = false
		}
	}
}

func choose(src *[]instr) string {
	// look at the instructions
	// pick the first pre according to alphabet
	m := make(map[string]int, 0)
	for i := range *src {
		_tmp := (*src)[i]
		m[_tmp.pre]++
	}
	fmt.Println("choices in choose", m)
	// need to sort now
	keys := make([]string, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	res := keys[0]
	markNotAvail(res, src)
	return res
}

func addToAvail(picked string, src *[]instr, avail *[]instr) {
	// we just picked one
	// pull the anything from src mark not avail
	// that has a matching pre and move to avail
}

func part1wrong() {
	filebytes := readfile("example-input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")
	var instructions = make([]instr, 0)
	for i := range filelines {
		line := filelines[i]
		fmt.Println(line)
		var pre, post string
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &pre, &post)
		fmt.Println("\t", pre, post)
		instructions = append(instructions, instr{pre, post, true})
	}
	var avail = make([]instr, 0)
	movePre(instructions[0].pre, &avail, &instructions)
	fmt.Println("after move instructions", instructions)
	fmt.Println("after move avail", avail)

	res := make([]string, 0)
	picked := choose(&avail)
	fmt.Println("we picked", picked)
	fmt.Println("res so far...", res)
	fmt.Println("avail", avail)
	fmt.Println("instructions", instructions)
	res = append(res, picked)
	addToAvail(picked, &instructions, &avail)
}

func part1() {
	filebytes := readfile("example-input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")
	var instructions = make([]instr, 0)
	for i := range filelines {
		line := filelines[i]
		fmt.Println(line)
		var pre, post string
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &pre, &post)
		fmt.Println("\t", pre, post)
		instructions = append(instructions, instr{pre, post, true})
	}
	fmt.Println("all instructions", instructions)
	res := make([]string, 0)

	fmt.Println("Res", res)
}

func main() {
	part1()
}
