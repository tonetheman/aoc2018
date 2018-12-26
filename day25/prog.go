package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readfile(filename string) []byte {
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("could not read file")
		return nil
	}
	return filedata
}

type _point struct {
	x, y, z, m int
}

func dist(p1, p2 _point) int {
	tabs := func(v int) int {
		if v < 0 {
			return v * -1
		}
		return v
	}
	_x := p1.x - p2.x
	_x = tabs(_x)
	_y := p1.y - p2.y
	_y = tabs(_y)
	_z := p1.z - p2.z
	_z = tabs(_z)
	_m := p1.m - p2.m
	_m = tabs(_m)
	return _x + _y + _z + _m
}

// constellation is an array of points
type _cons struct {
	points []_point
}

// galaxy is an array of constellations
type _gal struct {
	cons []_cons
}

func part1() {
	//filebytes := readfile("example1")
	filebytes := readfile("example3")

	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")
	points := make([]_point, 0)
	for i := range filelines {
		p := _point{}
		line := filelines[i]
		fmt.Sscanf(line, "%d,%d,%d,%d", &p.x, &p.y, &p.z, &p.m)
		fmt.Println(p, dist(_point{}, p))
		points = append(points, p)
	}

	var g _gal
	// try to find a home for a point
	for i := range points {
		placed := false
		p := points[i] // current point we need to place
		// look in each constellation in the galaxy
		for j := range g.cons {
			cc := g.cons[j]
			for k := range cc.points {
				d := dist(p, cc.points[k])
				if d <= 3 {
					placed = true
					//fmt.Println("added to cc.points")
					g.cons[j].points = append(g.cons[j].points, p)
				}
			}
		}
		// if not placed this goes to a new constellation
		if !placed {
			fmt.Println("point not placed", p)
			cc := _cons{}
			cc.points = make([]_point, 1)
			cc.points[0] = p
			g.cons = append(g.cons, cc)
		}
		fmt.Println("galaxy at end of searching", g)
	}
	fmt.Println("num of cons in galaxy is", len(g.cons))
}

func main() {
	part1()
}
