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
	var m map[byte]int

	return m
}

func (g *_grid) process() {
	for i := 0; i < _size; i++ {
		for j := 0; j < _size; j++ {
			val := g.get(i, j)
			sres := g.survey(i, j)
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
