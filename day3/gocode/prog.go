package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"text/scanner"
)

func readfile(filename string) []byte {
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("could not read file")
		return nil
	}
	return filedata
}

func bytesToString(buffer []byte) string {
	return string(buffer)
}

func scan(str string) {
	var sc scanner.Scanner
	sc.Init(strings.NewReader(str))
	count := 0
	for tok := sc.Scan(); tok != scanner.EOF; tok = sc.Scan() {
		fmt.Println("%s %s\n", sc.Position, sc.TokenText())
		count++
		if count > 9 {
			break
		}
	}
}

func main() {
	bytedata := readfile("../input")
	stringdata := bytesToString(bytedata)
	scan(stringdata)
}
