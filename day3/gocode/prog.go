package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"runtime"
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

func markboardbad(id int, sq *[esize][esize]int) {
	for i := 0; i < esize; i++ {
		for j := 0; j < esize; j++ {
			if sq[j][i] == id {
				sq[j][i] = -1
			}
		}
	}
}

func layoutoutonepart2(c cloth, sq *[esize][esize]int) bool {
	bad := false
	badids := make(map[int]int, 0)
	for i := c.posx; i < c.posx+c.width; i++ {
		for j := c.posy; j < c.posy+c.height; j++ {
			if sq[j][i] != 0 {
				bad = true
				if sq[j][i] != -1 {
					badids[sq[j][i]]++
				}
			}
			//sq[j][i] = c.id
		}
	}
	if bad {
		// there something down already underneath
		for i := c.posx; i < c.posx+c.width; i++ {
			for j := c.posy; j < c.posy+c.height; j++ {
				sq[j][i] = -1
			}
		}
		// mark the bad ones too
		for k, _ := range badids {
			markboardbad(k, sq)
		}
	} else {
		for i := c.posx; i < c.posx+c.width; i++ {
			for j := c.posy; j < c.posy+c.height; j++ {
				sq[j][i] = c.id
			}
		}
	}
	return bad
}

func layout(cloths []cloth, sq *[esize][esize]int) {
	for i := 0; i < len(cloths); i++ {
		layoutone(cloths[i], sq)
	}
}

func layoutpart2(cloths []cloth, sq *[esize][esize]int) {
	for i := 0; i < len(cloths); i++ {
		res := layoutoutonepart2(cloths[i], sq)
		fmt.Println(res)
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

func examplepart2() {
	fmt.Println("starting example part 2")
	var sq [esize][esize]int
	s := `
	#1 @ 1,3: 4x4
	#2 @ 3,1: 4x4
	#3 @ 5,5: 2x2
	`
	sdata := rt(s)
	cloths := parseall(sdata)
	layoutpart2(cloths, &sq)
	//pr(sq)

	validid := -1
	for i := 0; i < esize; i++ {
		for j := 0; j < esize; j++ {
			if sq[i][j] > 0 {
				fmt.Println(sq[i][j])
				validid = sq[i][j]
			}
		}
	}
	fmt.Println("answer", validid)
}

func simpleparseall(sdata []string) []cloth {
	res := make([]cloth, 0)
	for i := 0; i < len(sdata); i++ {
		var id, posx, posy, w, h int
		fmt.Sscanf(sdata[i], "#%d @ %d,%d: %dx%d", &id, &posx, &posy, &w, &h)
		fmt.Println(id, posx, posy, w, h)
		res = append(res, cloth{id, posx, posy, w, h})
	}
	return res
}

func part2() {
	// get bytes
	bytedata := readfile("../input")
	// fix windows crap 1310
	if runtime.GOOS == "windows" {
		// stupid idea to fix problems with parsing ...
		// did not appear to work
		for i := 0; i < len(bytedata); i++ {
			if bytedata[i] == '\r' {
				bytedata[i] = ' '
			}
		}
	}
	// convert data to string
	stringdata := string(bytedata)
	// split on the \n
	data := strings.Split(stringdata, "\n")
	// much easy parse i am an eeeeeediot
	cloths := simpleparseall(data)
	//fmt.Println("part2", cloths)
	var sq [esize][esize]int
	layoutpart2(cloths, &sq)
	validid := -1
	//pr(sq)
	for i := 0; i < esize; i++ {
		for j := 0; j < esize; j++ {
			if sq[i][j] > 0 {
				fmt.Println(sq[i][j])
				validid = sq[i][j]
			}
		}
	}
	fmt.Println("answer", validid)
}

func main() {
	//part1()
	//examplepart2()
	part2()
}
