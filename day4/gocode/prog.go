package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type _rec struct {
	y, m, d, hh, mm int
	desc            string
}

type byTimestamp []_rec

func (a byTimestamp) Len() int {
	return len(a)
}
func (a byTimestamp) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a byTimestamp) Less(i, j int) bool {
	if a[i].y != a[j].y {
		return a[i].y < a[j].y
	}
	if a[i].m != a[j].m {
		return a[i].m < a[j].m
	}
	if a[i].d != a[j].d {
		return a[i].d < a[j].d
	}
	if a[i].hh != a[j].hh {
		return a[i].hh < a[j].hh
	}
	return a[i].mm < a[j].mm
}

func readfile(filename string) []byte {
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("could not read file")
		return nil
	}
	return filedata
}

func getInputRecords(filename string) []_rec {
	filebytes := readfile(filename)
	// fix windows/dos
	for i := 0; i < len(filebytes); i++ {
		if filebytes[i] == '\r' {
			filebytes[i] = ' '
		}
	}
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")

	_recs := make([]_rec, 0)
	for i := 0; i < len(filelines); i++ {
		var r _rec
		// this will get timestamp on front
		fmt.Sscanf(filelines[i], "[%d-%d-%d %d:%d]",
			&r.y, &r.m, &r.d, &r.hh, &r.mm)

		_tmplinedata := strings.Split(filelines[i], "]")
		r.desc = strings.Trim(_tmplinedata[1], " ")
		//fmt.Println(filelines[i])
		//fmt.Println(r)
		_recs = append(_recs, _rec{r.y, r.m, r.d, r.hh, r.mm, r.desc})
	}
	return _recs
}

func pr(recs []_rec) {
	for i := 0; i < len(recs); i++ {
		fmt.Println(i, recs[i])
	}
}

func pr10(recs []_rec) {
	for i := 0; i < 10; i++ {
		fmt.Println(recs[i])
	}
}

func validate_hour(recs []_rec) {
	var m map[int]int = make(map[int]int)
	for i := 0; i < len(recs); i++ {
		r := recs[i]
		m[r.hh]++
	}
	fmt.Println(m)
	// PROVED: hours only 23 and 0 in the file for real input
	// and example input
}

func main() {
	//recs := getInputRecords("../input")
	recs := getInputRecords("./input-example")
	junk := byTimestamp(recs)
	sort.Sort(junk)
	pr(recs)

}
