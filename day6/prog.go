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

func main() {
	filebytes := readfile("example-input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")
	fmt.Println(filelines)
}
