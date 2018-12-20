package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const _size = 10

type _grid [_size * _size]byte
type _p struct{ offsetrow, offsetcol int }

func readfile(filename string) []byte {
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("could not read file")
		return nil
	}
	return filedata
}

func (g *_grid) pr() {
	for i := 0; i < _size; i++ {
		for j := 0; j < _size; j++ {
			fmt.Print(g[i*_size+j])
		}
		fmt.Println()
	}
}
func (g *_grid) set(row, col int, val byte) {
	g[row*_size+col] = val
}
func (g *_grid) get(row, col int) byte {
	return g[row*_size+col]
}
func (g *_grid) survey(row, col int) map[byte]int {
	debug := false
	if row == 0 && col == 0 {
		debug = true
	}
	m := make(map[byte]int)
	offsets := []_p{{-1, -1}, {-1, 0}, {-1, 1}}
	for i := range offsets {
		o := offsets[i]
		if row+o.offsetrow >= 0 && row+o.offsetrow < _size {
			if col+o.offsetcol >= 0 && col+o.offsetcol < _size {
				if debug {
					fmt.Println(row, col, row+o.offsetrow, col+o.offsetcol)
				}
				m[g.get(row+o.offsetrow, col+o.offsetcol)]++
			}
		}
	}
	return m
}

func (g *_grid) process() {
	for i := 0; i < _size; i++ {
		for j := 0; j < _size; j++ {
			val := g.get(i, j)
			sres := g.survey(i, j)
			if i == 0 && j == 0 {
				fmt.Println(sres)
			}
			if val == 1 { // open
				if sres[1] >= 3 {
					// change
				}
			}
		}
	}
}
func main() {
	filebytes := readfile("example-input")
	filestring := string(filebytes)
	fmt.Println(filestring)
	filelines := strings.Split(filestring, "\n")
	var g _grid
	for i := range filelines {
		for j := range filelines[i] {
			s := filelines[i]
			cc := s[j]
			if cc == '.' {
				//fmt.Println("grnd")
				//g[i*_size+j] = 1
				g.set(i, j, 1)
			}
			if cc == '|' {
				//g[i*_size+j] = 2
				g.set(i, j, 2)
			}
			if cc == '#' {
				//g[i*_size+j] = 3
				g.set(i, j, 3)
			}
		}
	}
	g.pr()
	g.process()
}
