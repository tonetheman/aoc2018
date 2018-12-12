package main

import (
	"fmt"
	"io/ioutil"
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

type _rule struct {
	_pat [5]rune
	_res rune
}

const SS = 100
const zeroOffset = 50

func createRule(s string) _rule {
	var tmp _rule
	tmp._pat[0] = rune(s[0])
	tmp._pat[1] = rune(s[1])
	tmp._pat[2] = rune(s[2])
	tmp._pat[3] = rune(s[3])
	tmp._pat[4] = rune(s[4])
	tmp._res = rune(s[9])
	//for pos, val := range s {
	//	tmp._pat[pos] = val
	//}
	return tmp
}

func matchrule(r _rule, pos int, state [SS]rune) bool {
	//fmt.Println("matchrule", pos, pos+4)
	res := r._pat[0] == state[pos] &&
		r._pat[1] == state[pos+1] &&
		r._pat[2] == state[pos+2] &&
		r._pat[3] == state[pos+3] &&
		r._pat[4] == state[pos+4]
	return res
}

func printState(state [SS]rune) {
	for i := 0; i < SS; i++ {
		if i == zeroOffset {
			fmt.Print("0")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	for i := 0; i < SS; i++ {
		fmt.Printf("%c", state[i])
	}
	fmt.Println()
}

func example() {

	filebytes := readfile("example-input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")
	fmt.Println(filelines)
	fmt.Println()
	fmt.Println(filelines[0])
	initialStateString := strings.Split(filelines[0], ": ")[1]
	var state [SS]rune
	for i := 0; i < SS; i++ {
		state[i] = rune('.')
	}
	for pos, val := range initialStateString {
		fmt.Println(pos, val)
		state[zeroOffset+pos] = val
	}
	// state we will keep to work on
	printState(state)
	// parse rules now
	// filelines[2] is start filelines[1] empty
	rules := make([]_rule, 0)
	for i := 2; i < len(filelines); i++ {
		rules = append(rules, createRule(filelines[i]))
	}
	//fmt.Println(rules)

	// try to match rule[0]
	// -4 so we can hit the full range of the state
	//for i := 0; i < SS-4; i++ {
	//	if matchrule(rules[0], i, state) {
	//		fmt.Println("matched", i-zeroOffset)
	//	}
	//}
	var nextstate [100]rune
	for i := 0; i < SS; i++ {
		nextstate[i] = rune('.')
	}

	for i := 0; i < SS-4; i++ {

		// for each pos in state
		// check all the rules? or just match one
		for j := 0; j < len(rules); j++ {
			if matchrule(rules[j], i, state) {
				// setup next state
				nextstate[i] = rules[i]._res
			}
		}
	}

	printState(nextstate)
}

func main() {
	example()
}
