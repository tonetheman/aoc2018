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
	ilen := len(input) // take this here
	// create an empty array
	// used for clearing in go with a copy (sneaky)
	emptyarray := make([]byte, ilen)
	//fmt.Println("part1 ilen is", ilen)
	tmpinput := make([]byte, len(input))
	cleartmpinput := func() {
		// slowest
		//for i := 0; i < len(input); i++ {
		//	tmpinput[i] = 0
		//}
		// faster
		//for i := range input {
		//	tmpinput[i] = 0
		//}
		// fastest!!!
		copy(tmpinput, emptyarray)
	}
	movetotmp := func() {
		counter := 0
		for i := 0; i < ilen; i++ {
			if input[i] != 0 {
				tmpinput[counter] = input[i]
				counter++
			}
		}
	}
	clearinput := func() {
		copy(input, emptyarray)
		//for i := 0; i < ilen; i++ {
		//	input[i] = 0
		//}
	}
	movetoinput := func() {
		// slow compared to copy
		//for i := 0; i < ilen; i++ {
		//	input[i] = tmpinput[i]
		//}
		copy(input, tmpinput)
	}

	for {
		cleartmpinput()
		gotone := false

		//fmt.Println("input before cycle", string(input))
		for i := 0; i < ilen-1; i++ {
			c1 := input[i]
			c2 := input[i+1]
			if c1 == c2+32 || c1+32 == c2 {
				//fmt.Println("got one", i, c1, c2)
				input[i] = 0
				input[i+1] = 0
				gotone = true
				//break
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

		if gotone {
			ilen -= 2
		}

		if !gotone {
			//fmt.Println("QUITING")
			break
		}

	}

	// clear all the input here not trimmed down amount
	// being careful
	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			input[i] = 32
		}
	}
	tmp := string(input)
	tmp2 := strings.Trim(tmp, " ")
	//fmt.Println("ENDCOMPRESS", tmp2)
	//fmt.Println(len(tmp2))
	//fmt.Println(len(tmp2))
	return len(tmp2)
}

func removeP(c byte, input []byte) []byte {
	// remove a specific item and compress
	// then return it
	size := len(input)
	for i := 0; i < len(input); i++ {
		if input[i] == c || input[i] == c-32 {
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

func part2() {
	input := readfile("./input")
	//input := readfile("./input-example")
	//fmt.Println("original input here", input)
	//fmt.Println(string(input))
	//originputlen := len(input)
	savedinput := make([]byte, len(input))
	//for i := 0; i < len(input); i++ {
	//	savedinput[i] = input[i]
	//}
	copy(savedinput, input)
	restoreinput := func() {
		//for i := 0; i < originputlen; i++ {
		//	input[i] = savedinput[i]
		//}
		copy(input, savedinput)
	}
	smallest := 90000000
	mychars := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	//mychars := []byte{'c'}
	for ii := 0; ii < len(mychars); ii++ {
		restoreinput()
		cc := mychars[ii]
		newinput := removeP(cc, input)
		//fmt.Println(string(input))
		//fmt.Println(string(newinput))
		res := part1(newinput)
		if res < smallest {
			smallest = res
		}
		//fmt.Println("compressed", cc, part1(newinput))
	}
	fmt.Println("smallest is", smallest)
}

func main() {
	part2()
}
