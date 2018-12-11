package main

import (
	"errors"
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

func part1wrongagain() {
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

type _HeadPointer struct {
	count    int
	name     string
	children []string
}

func findHeadPointer(val string, heads []_HeadPointer) (*_HeadPointer, error) {
	for i := range heads {
		h := heads[i]
		if h.name == val {
			return &h, nil
		}
	}
	return &_HeadPointer{}, errors.New("could not find")
}

func replacePointer(src _HeadPointer, heads *[]_HeadPointer) {
	for i := range *heads {
		h := (*heads)[i]
		if h.name == src.name {
			(*heads)[i] = src
		}
	}
}

func getinstructions() []instr {
	//filebytes := readfile("example-input")
	filebytes := readfile("input")
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
	fmt.Println()
	return instructions
}

func part1() {
	instructions := getinstructions()

	heads := make([]_HeadPointer, 0)
	count := 0
	for {

		_instr := instructions[count]
		fmt.Println("processing this instr", _instr)

		_ptr, err := findHeadPointer(_instr.pre, heads)
		if err != nil {
			// need to insert a records for pre
			_tmp := _HeadPointer{}
			_tmp.name = _instr.pre
			_tmp.children = make([]string, 0)
			_tmp.children = append(_tmp.children, _instr.post)
			heads = append(heads, _tmp)
		} else {
			fmt.Println("adding to children", _ptr)
			fmt.Printf("memaddr children %p\n", &_ptr.children)
			_ptr.children = append(_ptr.children, _instr.post)
			fmt.Printf("memaddr children after append %p\n", _ptr.children)
			fmt.Println("after adding to children", _ptr)
			replacePointer(*_ptr, &heads)
		}

		count++
		if count >= len(instructions) {
			break
		}
	}

	// now fill out counts
	// fill out zeros first
	degreemap := make(map[string]int)
	for i := range heads {
		degreemap[heads[i].name] = 0
	}
	// now do counts
	for i := range heads {
		h := heads[i]
		for j := range h.children {
			c := h.children[j]
			degreemap[c]++
		}
	}
	fmt.Println("built graph...")
	fmt.Println("heads", heads)
	fmt.Println("degreemap", degreemap)
	fmt.Println()

	// now find all zero in degree map
	count = 0
	res := make([]string, 0)
	for {
		// make a new array to hold the choices
		// we can get more than one
		// and in that case we need to use lexographic sort
		holdszeros := make([]string, 0)
		for k, v := range degreemap {
			if v == 0 {
				holdszeros = append(holdszeros, k)
			}
		}
		if len(holdszeros) == 0 {
			fmt.Println("WARNING: no choices!")
			break
		}
		// sort the choices
		sort.Strings(holdszeros)
		fmt.Println(count, "choices", holdszeros)
		// the choice for this round is the guy in index 0
		// mark him out of the game
		degreemap[holdszeros[0]] = -1
		// put him in the result
		res = append(res, holdszeros[0])
		// get pointer, needed to mark children
		_ptr, ok := findHeadPointer(holdszeros[0], heads)
		if ok != nil {
			fmt.Println("ERR: could not find pointer")
			break
		}
		// reduce the degree of each child in the degree map
		for i := range _ptr.children {
			c := _ptr.children[i]
			degreemap[c]--
		}

		// assume we need to stop
		stop := true
		for _, v := range degreemap {
			// if anything in degree map has not been
			// marked we need to keep going
			if v != -1 {
				stop = false
			}
		}

		count++
		if stop {
			break
		}
	}
	fmt.Println("made choice...")
	fmt.Println("heads", heads)
	fmt.Println("degreemap", degreemap)
	fmt.Println("res", res)
	fmt.Println(strings.Join(res, ""))
	fmt.Println()

}

type _worker struct {
	choice string
	count  int
	busy   bool
}

func findNonBusyWorker(w []_worker) int {
	for i := range w {
		if w[i].busy == false {
			return i
		}
	}
	return -1
}

func part2() {
	instructions := getinstructions()

	heads := make([]_HeadPointer, 0)
	count := 0
	for {

		_instr := instructions[count]
		fmt.Println("processing this instr", _instr)

		_ptr, err := findHeadPointer(_instr.pre, heads)
		if err != nil {
			// need to insert a records for pre
			_tmp := _HeadPointer{}
			_tmp.name = _instr.pre
			_tmp.children = make([]string, 0)
			_tmp.children = append(_tmp.children, _instr.post)
			heads = append(heads, _tmp)
		} else {
			fmt.Println("adding to children", _ptr)
			fmt.Printf("memaddr children %p\n", &_ptr.children)
			_ptr.children = append(_ptr.children, _instr.post)
			fmt.Printf("memaddr children after append %p\n", _ptr.children)
			fmt.Println("after adding to children", _ptr)
			replacePointer(*_ptr, &heads)
		}

		count++
		if count >= len(instructions) {
			break
		}
	}

	// now fill out counts
	// fill out zeros first
	degreemap := make(map[string]int)
	for i := range heads {
		degreemap[heads[i].name] = 0
	}
	// now do counts
	for i := range heads {
		h := heads[i]
		for j := range h.children {
			c := h.children[j]
			degreemap[c]++
		}
	}
	fmt.Println("built graph...")
	fmt.Println("heads", heads)
	fmt.Println("degreemap", degreemap)
	fmt.Println()

	// current time
	second := 0
	workers := make([]_worker, 5)
	// now find all zero in degree map
	count = 0
	res := make([]string, 0)
	for {
		// make a new array to hold the choices
		// we can get more than one
		// and in that case we need to use lexographic sort
		holdszeros := make([]string, 0)
		for k, v := range degreemap {
			if v == 0 {
				holdszeros = append(holdszeros, k)
			}
		}
		if len(holdszeros) == 0 {
			fmt.Println("WARNING: no choices!")
			break
		}
		// sort the choices
		sort.Strings(holdszeros)
		fmt.Println(count, "choices", holdszeros)
		// the choice for this round is the guy in index 0
		// mark him out of the game
		degreemap[holdszeros[0]] = -1
		// put him in the result
		res = append(res, holdszeros[0])
		// get pointer, needed to mark children
		_ptr, ok := findHeadPointer(holdszeros[0], heads)
		if ok != nil {
			fmt.Println("ERR: could not find pointer")
			break
		}
		// reduce the degree of each child in the degree map
		for i := range _ptr.children {
			c := _ptr.children[i]
			degreemap[c]--
		}

		// assume we need to stop
		stop := true
		for _, v := range degreemap {
			// if anything in degree map has not been
			// marked we need to keep going
			if v != -1 {
				stop = false
			}
		}

		count++
		if stop {
			break
		}
	}
	fmt.Println("Second", second)
	fmt.Println("workers", workers)
	fmt.Println("made choice...")
	fmt.Println("heads", heads)
	fmt.Println("degreemap", degreemap)
	fmt.Println("res", res)
	fmt.Println(strings.Join(res, ""))
	fmt.Println()
}

func main() {
	part1()
}
