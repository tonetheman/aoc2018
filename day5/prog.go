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

func part1(input []byte) int {
	tmpinput := make([]byte, len(input))
	cleartmpinput := func() {
		for i := 0; i < len(input); i++ {
			tmpinput[i] = 0
		}
	}
	movetotmp := func() {
		counter := 0
		for i := 0; i < len(input); i++ {
			if input[i] != 0 {
				tmpinput[counter] = input[i]
				counter++
			}
		}
	}
	clearinput := func() {
		for i := 0; i < len(input); i++ {
			input[i] = 0
		}
	}
	movetoinput := func() {
		for i := 0; i < len(input); i++ {
			input[i] = tmpinput[i]
		}
	}

	for {
		cleartmpinput()
		gotone := false

		//fmt.Println("input before cycle", string(input))
		for i := 0; i < len(input)-1; i++ {
			c1 := input[i]
			c2 := input[i+1]
			if c1 == c2+32 || c1+32 == c2 {
				//fmt.Println("got one", i, c1, c2)
				input[i] = 0
				input[i+1] = 0
				gotone = true
				break
			}
		}
		//fmt.Println("input after clear", string(input))
		// compress it
		movetotmp()
		// clear it
		clearinput()
		// put it back
		movetoinput()
		// need to compress gaps here

		if !gotone {
			fmt.Println("QUITING")
			break
		}

	}

	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			input[i] = 32
		}
	}
	tmp := string(input)
	tmp2 := strings.Trim(tmp, " ")
	//fmt.Println(tmp2)
	//fmt.Println(len(tmp2))
	fmt.Println(len(tmp2))
	return len(tmp2)
}

func removeP(c byte, input []byte) []byte {
	// remove a specific item and compress
	// then return it
	size := len(input)
	for i := 0; i < len(input); i++ {
		if input[i] == c || input[i] == c+32 {
			input[i] = 0
			size--
		}
	}

	var tmpinput []byte = make([]byte, size)

	// now compress it
	counter := 0
	for i := 0; i < len(input); i++ {
		if input[i] != 0 {
			tmpinput[counter] = input[i]
			counter++
		}
	}
	return tmpinput
}

func main() {
	//input := readfile("./input")
	input := readfile("./input-example")
	newinput := removeP('a', input)
	fmt.Println(string(input))
	fmt.Println(string(newinput))
	part1(input)
}
