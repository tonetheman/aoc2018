package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type _board struct {
	grid  []string
	width int
}
type _loc struct {
	row int
	col int
}

func readfile(filename string) []byte {
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("could not read file")
		return nil
	}
	return filedata
}

func (b _board) getWidth() int {
	s := b.grid[0]
	return len(s)
}

func (b _board) findAllInRange(row, col int) []_loc {
	player := b.grid[row][col]
	fmt.Printf("findAllInRange for %d %d %c\n", row, col, player)
	res := make([]_loc, 0)
	for i := range b.grid {
		for j := 0; j < b.width; j++ {
			if b.grid[i][j] == 'G' {
				res = append(res, _loc{i, j})
			}
		}
	}
	return res
}

func main() {
	filebytes := readfile("example2_input")
	filestring := string(filebytes)
	b := _board{}
	b.grid = strings.Split(filestring, "\n")
	for i := range b.grid {
		b.grid[i] = strings.Trim(b.grid[i], " \r\n")
		fmt.Println(i, b.grid[i])
	}
	fmt.Println("width", b.getWidth())
	b.width = b.getWidth()
	inrange := b.findAllInRange(1, 1)
	fmt.Println("in range enemies", inrange)
}
