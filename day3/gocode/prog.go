package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type cloth struct {
	id            int
	posx, posy    int
	width, height int
}

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
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)
	for i := 0; i < 4; i++ {
		fmt.Println(scanner.Scan())
		fmt.Println(scanner.Text())
	}
}

func regexTest(str string) {
	P := regexp.MustCompile("^#([0-9])+ @ [0-9]+,[0-9]+: [0-9]+x[0-9]+")
	data := P.Split(str, -1)
	fmt.Println(data)
}

func main() {
	bytedata := readfile("../input")
	stringdata := bytesToString(bytedata)
	regexTest(stringdata)
}
