package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

func dist3(x1, y1, z1, x2, y2, z2 int) int {
	return int(math.Abs(float64(x1-x2)) +
		math.Abs(float64(y1-y2)) +
		math.Abs(float64(z1-z2)))
}

func part1() {
	//filebytes := readfile("example-input")
	filebytes := readfile("input")

	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")

	// largest radius index and value
	lr_index := -1
	lr_value := -10000

	for i := range filelines {
		line := filelines[i]
		var x, y, z, r int
		fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d",
			&x, &y, &z, &r)
		fmt.Println(x, y, z, r)

		// find largest radius
		if r > lr_value {
			lr_index = i
			lr_value = r
		}
	}

	fmt.Println("lr_index", lr_index)
	fmt.Println("lr_value", lr_value)
	var lrx, lry, lrz, lrr int
	fmt.Sscanf(filelines[lr_index],
		"pos=<%d,%d,%d>, r=%d",
		&lrx, &lry, &lrz, &lrr)

	in_range_count := 0
	for i := range filelines {
		var x, y, z, r int
		fmt.Sscanf(filelines[i],
			"pos=<%d,%d,%d>, r=%d",
			&x, &y, &z, &r)
		d := dist3(lrx, lry, lrz, x, y, z)
		//fmt.Println("dist calc", i, d)
		if d <= lrr {
			in_range_count++
		}
	}
	fmt.Println("in range count", in_range_count)

}

func part2() {
	//filebytes := readfile("example2-input")
	filebytes := readfile("input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")
	//fmt.Println(filelines)
	// find boundaries
	mx := -1
	my := -1
	mz := -1
	for i := range filelines {
		var x, y, z, r int
		fmt.Sscanf(filelines[i],
			"pos=<%d,%d,%d>, r=%d",
			&x, &y, &z, &r)
		//fmt.Println(x, y, z)
		if x > mx {
			mx = x
		}
		if y > my {
			my = y
		}
		if z > mz {
			mz = z
		}
	}
	fmt.Println("largest", mx, my, mz)

	type _point struct {
		x, y, z, r int
	}
	var points = make([]_point, 0)
	for i := range filelines {
		var x, y, z, r int
		fmt.Sscanf(filelines[i],
			"pos=<%d,%d,%d>, r=%d",
			&x, &y, &z, &r)
		points = append(points, _point{x, y, z, r})
	}
	// we know index 938 is the guy
	var lx, ly, lz, lr int
	fmt.Sscanf(filelines[938],
		"pos=<%d,%d,%d>, r=%d",
		&lx, &ly, &lz, &lr)
	fmt.Println("dude", lx, ly, lz, lr)
	for i := lx - lr; i <= lx+lr; i++ {
		for j := ly - lr; j <= ly+lr; j++ {
			for k := lz - lr; k <= lz+lr; k++ {

				// look at every point
				// determine the number of bots in
				// range of this point
				in_range := 0
				for ii := range points {
					p := points[ii]
					d := dist3(i, j, k, p.x, p.y, p.z)
					//fmt.Println(i, j, k, r, d, (d <= r))
					if d <= p.r {
						in_range++
					}
				}
				if in_range > 3 {
					fmt.Println("!!!!this point",
						i, j, k, in_range)
				}

			}
		}
	}
}

func main() {
	//part1()
	part2()
}
