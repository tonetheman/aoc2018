package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// example size board
//const esize = 10

const esize = 1000

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
	height, _ := strconv.Atoi(strings.TrimRight(_whd[1], "\n"))
	return cloth{id, posx, posy, width, height}
}

func parseall(s []string) []cloth {
	res := make([]cloth, 0)
	for i := 0; i < len(s); i++ {
		res = append(res, parseone(s[i]))
	}
	return res
}

func pr(sq [esize][esize]int) {
	for i := 0; i < esize; i++ {
		fmt.Println(sq[i])
	}
}

func layoutone(c cloth, sq *[esize][esize]int) {
	fmt.Println("layout 1", c)
	for i := c.posx; i < c.posx+c.width; i++ {
		for j := c.posy; j < c.posy+c.height; j++ {
			sq[j][i]++
		}
	}
}

func layout(cloths []cloth, sq *[esize][esize]int) {
	for i := 0; i < len(cloths); i++ {
		layoutone(cloths[i], sq)
	}
}

func example() {
	var sq [esize][esize]int
	pr(sq)
	s := `
	#1 @ 1,3: 4x4
	#2 @ 3,1: 4x4
	#3 @ 5,5: 2x2
	`
	sdata := rt(s)
	cloths := parseall(sdata)
	layout(cloths, &sq)
	count := 0
	for i := 0; i < esize; i++ {
		for j := 0; j < esize; j++ {
			if sq[i][j] > 1 {
				count++
			}
		}
	}
	fmt.Println("answer", count)
}

func part1() {
	bytedata := readfile("../input")
	stringdata := bytesToString(bytedata)
	data := rt(stringdata)
	cloths := parseall(data)
	var sq [esize][esize]int
	layout(cloths, &sq)
	count := 0
	for i := 0; i < esize; i++ {
		for j := 0; j < esize; j++ {
			if sq[i][j] > 1 {
				count++
			}
		}
	}
	fmt.Println("answer", count)
}

func main() {
	part1()
}
