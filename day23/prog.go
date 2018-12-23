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
		fmt.Println("dist calc", i, d)
		if d <= lrr {
			in_range_count++
		}
	}
	fmt.Println("in range count", in_range_count)

}

func main() {
	part1()

}
