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

func part1() {
	filebytes := readfile("example1")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")
	for i := range filelines {
		p := _point{}
		line := filelines[i]
		fmt.Sscanf(line, "%d,%d,%d,%d", &p.x, &p.y, &p.z, &p.m)
		fmt.Println(p, dist(_point{}, p))
	}
}

func main() {
	part1()
}
