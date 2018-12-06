package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

const GS = 360

type grid [GS][GS]int

func (g *grid) set_grid(row, col int, val int) {
	g[col][row] = val
}
func (g *grid) get_grid(row, col int) int {
	return g[col][row]
}

type point struct {
	id       int
	row, col int
}
type points []point

func readfile(filename string) []byte {
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("could not read file")
		return nil
	}
	return filedata
}

func pr(g *grid) {
	names := []string{".", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for i := 0; i < GS; i++ {
		for j := 0; j < GS; j++ {
			fmt.Printf("%s ", names[g.get_grid(j, i)])
		}
		fmt.Println()
	}
}

func dist(p1row, p1col, p2row, p2col int) int {
	return int(math.Abs(float64(p1row-p2row)) + math.Abs(float64(p1col-p2col)))
}

func makePoints(filelines []string) points {
	res := make(points, 0)
	for i := range filelines {
		line := filelines[i]
		var row, col int
		fmt.Sscanf(line, "%d,%d", &row, &col)
		res = append(res, point{i + 1, row, col})
	}
	return res
}

func placePoints(scarypoints points, g *grid) {
	for i := range scarypoints {
		p := scarypoints[i]
		g.set_grid(p.row, p.col, p.id)
	}
}

func findWinner(row, col int, g *grid, scarypoints points) int {
	// determine distance for each scary point from
	// this point i,j - who is the winner of the point?
	m := make(map[int]int, 0)
	for i := range scarypoints {
		sc := scarypoints[i]
		pdist := dist(row, col, sc.row, sc.col)
		m[sc.id] = pdist
		//fmt.Println(sc, pdist, m)
	}
	fmt.Println(m)
	// now find the winner of the point
	smallest := 10000000
	smallest_count := 0
	smallest_id := -1
	for k, v := range m {
		if v < smallest {
			smallest = v
			smallest_id = k
		}
	}
	// determine how many won that point for ties
	matchCount := func() int {
		count := 0
		for _, v := range m {
			if v == smallest {
				count++
			}
		}
		return count
	}
	smallest_count = matchCount()

	// now you are done
	fmt.Println("small value", smallest)
	fmt.Println("small count", smallest_count)
	fmt.Println("small id", smallest_id)
	if smallest_count == 1 {
		return smallest_id
	} else {
		return -1
	}
}

func assignEmpty(scarypoints points, g *grid) {
	for i := 0; i < GS; i++ {
		for j := 0; j < GS; j++ {
			res := findWinner(j, i, g, scarypoints)
			if res == -1 {
				// tie
			} else {
				// not a tie
				g.set_grid(j, i, res)
			}

		}
	}
	pr(g)
	/*
		for i := 0; i < GS; i++ {
			for j := 0; j < GS; j++ {
				if g[i][j] == 0 {
					// empty point who is close?
					findWinner(i, j, g, scarypoints)
				}
			}
		}
	*/
}

func example() {
	filebytes := readfile("example-input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")
	scarypoints := makePoints(filelines)
	var g grid
	placePoints(scarypoints, &g)
	pr(&g)
	assignEmpty(scarypoints, &g)
}

func main() {

	filebytes := readfile("input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")
	scarypoints := makePoints(filelines)
	var g grid
	placePoints(scarypoints, &g)
	pr(&g)
	assignEmpty(scarypoints, &g)

	// max col 359 max row 349

}
