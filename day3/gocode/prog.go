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

// splits the string...
// FUCK regex in go this is easy in python
func regexTest(str string) []string {
	res := make([]string, 0)
	P := regexp.MustCompile("^#[0-9]+ @ [0-9]+,[0-9]+: [0-9]+x[0-9]+")
	firstone := P.FindString(str)
	res = append(res, firstone)
	//fmt.Println(firstone)
	data := P.Split(str, -1)
	//fmt.Println(data)
	for i := 0; i < len(data); i++ {
		res = append(res, data[i])
	}
	return res
}

func main() {
	bytedata := readfile("../input")
	stringdata := bytesToString(bytedata)
	splitData := regexTest(stringdata)
	fmt.Println(splitData)
}
