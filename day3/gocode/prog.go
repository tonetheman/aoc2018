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

//
func rt(str string) []string {
	P := regexp.MustCompile("#([0-9]+) @ [0-9]+,[0-9]+: [0-9]+x[0-9]+\n")
	all := P.FindAllString(str, -1) // this is all of the items
	return all
}

func main() {
	bytedata := readfile("../input")
	stringdata := bytesToString(bytedata)
	data := rt(stringdata)
	t := strings.Split(data[0], " ")
	fmt.Println(len(t), t)

	//splitData := regexTest(stringdata)
	//fmt.Println(splitData)
}
