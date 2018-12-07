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

func part1() {
	filebytes := readfile("example-input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")
	for i := range filelines {
		line := filelines[i]
		fmt.Println(line)
		var pre, post string
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &pre, &post)
		fmt.Println("\t", pre, post)
	}
}

func main() {
	part1()
}
