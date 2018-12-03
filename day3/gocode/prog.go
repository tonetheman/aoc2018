package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
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

func rt(str string) []string {
	P := regexp.MustCompile("#([0-9]+) @ [0-9]+,[0-9]+: [0-9]+x[0-9]+\n")
	all := P.FindAllString(str, -1) // this is all of the items
	return all
}

func parseone(s string) cloth {
	//fmt.Println("starting s", s)
	whdata := strings.Split(s, "#")
	withouthash := whdata[1]
	spdata := strings.Split(withouthash, " ")
	id, _ := strconv.Atoi(spdata[0])
	posxystring := spdata[2]
	widthheightstring := spdata[3]
	pdata := strings.Split(posxystring, ",")
	posx, _ := strconv.Atoi(pdata[0])
	posy, _ := strconv.Atoi(strings.TrimRight(pdata[1], ":"))
	_whd := strings.Split(widthheightstring, "x")
	width, _ := strconv.Atoi(_whd[0])
	height, _ := strconv.Atoi(_whd[1])
	return cloth{id, posx, posy, width, height}
}

func parseall(s []string) []cloth {
	res := make([]cloth, 0)
	for i := 0; i < len(s); i++ {
		res = append(res, parseone(s[i]))
	}
	return res
}

func main() {
	bytedata := readfile("../input")
	stringdata := bytesToString(bytedata)
	data := rt(stringdata)
	fmt.Println(parseall(data))
}
