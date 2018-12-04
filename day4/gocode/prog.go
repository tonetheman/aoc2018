package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type _rec struct {
	y, m, d, hh, mm int
	desc            string
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

func main() {
	getInputRecords("..\\input")
}
