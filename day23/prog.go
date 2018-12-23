package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
	"time"
)

func readfile(filename string) []byte {
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("could not read file")
		return nil
	}
	return filedata
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func dist3(x1, y1, z1, x2, y2, z2 int) int {
	//defer timeTrack(time.Now(), "dist3")
	//start := time.Now()
	_x := x1 + x2
	if _x < 0 {
		_x *= -1
	}
	_y := y1 + y2
	if _y < 0 {
		_y *= -1
	}
	_z := z1 + z2
	if _z < 0 {
		_z *= -1
	}
	//elapsed := time.Since(start)
	//log.Printf("Binomial took %s", elapsed)
	return _x + _y + _z
	/*
		return int(math.Abs(float64(x1-x2)) +
			math.Abs(float64(y1-y2)) +
			math.Abs(float64(z1-z2)))
	*/
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
	filebytes := readfile("example2-input")
	//filebytes := readfile("input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")
	//fmt.Println(filelines)

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
		fmt.Print(".")
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
				if in_range > 230 {
					fmt.Println("!!!!this point",
						i, j, k, in_range)
				}

			}
		}
	}
}

type _point struct {
	x, y, z, r int
}

func findMinMaxX(points []_point) (int, int) {
	ma := math.MinInt32
	mi := math.MaxInt32
	for i := range points {
		p := points[i]
		if p.x < mi {
			mi = p.x
		}
		if p.x > ma {
			ma = p.x
		}
	}
	return mi, ma
}

func findMinMaxY(points []_point) (int, int) {
	ma := math.MinInt32
	mi := math.MaxInt32
	for i := range points {
		p := points[i]
		if p.y < mi {
			mi = p.y
		}
		if p.y > ma {
			ma = p.y
		}
	}
	return mi, ma
}

func part2a() {
	filebytes := readfile("example2-input")
	//filebytes := readfile("input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")
	//fmt.Println(filelines)

	var points = make([]_point, 0)
	for i := range filelines {
		var x, y, z, r int
		fmt.Sscanf(filelines[i],
			"pos=<%d,%d,%d>, r=%d",
			&x, &y, &z, &r)
		points = append(points, _point{x, y, z, r})
	}

	// find list of points possible per x
	xc := 0
	distToZ := math.MaxInt32
	miX, maX := findMinMaxX(points)
	fmt.Println("mix and max", miX, maX)
	for i := miX; i <= maX; i++ {
		for ip := range points {
			p := points[ip]
			d := dist3(i, 0, 0, p.x, p.y, p.z)
			if d < p.r {
				//fmt.Println("*", i, 0, 0, d, p)
				xc++
				tmp := dist3(0, 0, 0, i, 0, 0)
				fmt.Println(i, 0, 0, tmp)
				if tmp < distToZ {
					distToZ = tmp
					fmt.Println("lowest so far", p, tmp)
				}
			}
		}
	}
	fmt.Println("xc is", xc)
	return
	// find choices for Y
	yc := 0
	miY, maY := findMinMaxY(points)
	for i := miY; i <= maY; i++ {
		for ip := range points {
			p := points[ip]
			d := dist3(0, i, 0, p.x, p.y, p.z)
			if d < p.r {
				//fmt.Println("*", i, 0, 0, d, p)
				yc++
			}
		}
	}

	fmt.Println("xc,yc", xc, yc)
}

func main() {
	//part1()
	part2a()
}
