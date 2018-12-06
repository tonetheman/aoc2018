package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

const GS = 20

type grid [GS][GS]int
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
	names := []string{".", "a", "b", "c", "d", "e", "f"}
	for i := 0; i < GS; i++ {
		for j := 0; j < GS; j++ {
			fmt.Printf("%s ", names[g[i][j]])
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
		g[p.col][p.row] = p.id
	}
}

func findWinner(col, row int, g *grid, scarypoints points) int {
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
		res := findWinner(i, 0, g, scarypoints)
		if res == -1 {
			// tie
		} else {
			// not a tie
			g[0][i] = res
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

func main() {
	filebytes := readfile("example-input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")
	scarypoints := makePoints(filelines)
	var g grid
	placePoints(scarypoints, &g)
	pr(&g)
	assignEmpty(scarypoints, &g)
}
