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

func dist3_test(x1, y1, z1, x2, y2, z2 int, out *int) {
	_x := x1 - x2
	if _x < 0 {
		_x *= -1
	}
	_y := y1 - y2
	if _y < 0 {
		_y *= -1
	}
	_z := z1 - z2
	if _z < 0 {
		_z *= -1
	}
	*out = _x + _y + _z
}

func dist3(x1, y1, z1, x2, y2, z2 int) int {
	//defer timeTrack(time.Now(), "dist3")
	//start := time.Now()

	_x := x1 - x2
	if _x < 0 {
		_x *= -1
	}
	_y := y1 - y2
	if _y < 0 {
		_y *= -1
	}
	_z := z1 - z2
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

// divide by 1,000,000
// this reduces the problem quite a bit
// but lets you search the full space
func part2_divideM() {
	/*
		best point found in this case
		50 49 26
		877 in range count
		0.114184 into search space
	*/
	//filebytes := readfile("example2-input")
	filebytes := readfile("input")
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

	F := 1000000
	for i := range points {
		p := &points[i]
		p.x /= F
		p.y /= F
		p.z /= F
		p.r /= F
	}
	//fmt.Println(points)
	/*
		junk := make([]int, 0)
		for i := range points {
			junk = append(junk, points[i].z)
		}
		sort.Ints(junk)
		//fmt.Println(junk)
		// full search space here
		//-144 187 x
		// -31 154 y
		// -71 127 z
	*/

	inRangeMax := 0
	totalSpace := 400 * 400 * 400
	count := 0
	for i := -147; i < 187+1; i++ {
		for j := -31; j < 154+1; j++ {
			for k := -71; k < 127+1; k++ {
				inRange := 0
				for ip := 0; ip < len(points); ip++ {
					p := &points[ip]
					d := dist3(i, j, k, p.x, p.y, p.z)
					if d < p.r {
						inRange++
					}
				}
				count++
				if inRange > inRangeMax {
					fmt.Println(i, j, k, inRange,
						float32(count)/float32(totalSpace))
					inRangeMax = inRange
				}
			}
		}
	}

}

// divide by 100,000
func part2_divideHT() {
	filebytes := readfile("input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")

	var points = make([]_point, 0)
	for i := range filelines {
		var x, y, z, r int
		fmt.Sscanf(filelines[i],
			"pos=<%d,%d,%d>, r=%d",
			&x, &y, &z, &r)
		points = append(points, _point{x, y, z, r})
	}

	F := 100000
	for i := range points {
		p := &points[i]
		p.x /= F
		p.y /= F
		p.z /= F
		p.r /= F
	}

	/*
			junk := make([]int, 0)
			for i := range points {
				junk = append(junk, points[i].z)
			}

		sort.Ints(junk)
		//fmt.Println(junk)
		// -1445 to 1871 x
		// -319 1544 y
		// -714 1278 z
		// from last time we know
		// we need something near this point
		// 50 49 26
	*/
	start := time.Now()
	count := 0
	inRangeMax := 0
	V := 100
	targetx := 0
	targety := 0
	targetz := 0
	lenPoints := len(points)
	for i := 500 - V; i < 500+V; i++ {
		for j := 490 - V; j < 490+V; j++ {
			for k := 260 - V; k < 260+V; k++ {

				inRange := 0
				for ip := 0; ip < lenPoints; ip++ {
					p := &points[ip]
					//d := dist3(i, j, k, p.x, p.y, p.z)
					var d int
					dist3_test(i, j, k, p.x, p.y, p.z, &d)
					if d < p.r {
						inRange++
					}
				}
				count++
				if inRange > inRangeMax {
					//fmt.Println(i, j, k, inRange)
					inRangeMax = inRange
					targetx = i
					targety = j
					targetz = k
				}

			}
		}
	}
	fmt.Println("inRange max", inRangeMax)
	fmt.Println(targetx, targety, targetz)
	elapsed := time.Since(start)
	fmt.Println("time", elapsed)
}

// divide by 10,000
func part2_divideTT() {
	/*
		from the 100,000 function we got this
		inRange max 892
		514 478 252
	*/
	filebytes := readfile("input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")

	var points = make([]_point, 0)
	for i := range filelines {
		var x, y, z, r int
		fmt.Sscanf(filelines[i],
			"pos=<%d,%d,%d>, r=%d",
			&x, &y, &z, &r)
		points = append(points, _point{x, y, z, r})
	}

	F := 10000
	for i := range points {
		p := &points[i]
		p.x /= F
		p.y /= F
		p.z /= F
		p.r /= F
	}

	/*
		junk := make([]int, 0)
		for i := range points {
			junk = append(junk, points[i].z)
		}

		sort.Ints(junk)
		fmt.Println(junk)
		// x  -14450 18712
		// y  -3192 15442
		// z  -7140 12783
	*/
	V := 100
	start := time.Now()
	count := 0
	inRangeMax := 0
	targetx := 0
	targety := 0
	targetz := 0
	lenPoints := len(points)

	for i := 5140 - V; i < 5140+V; i++ {
		for j := 4780 - V; j < 4780+V; j++ {
			for k := 2520 - V; k < 2520+V; k++ {

				inRange := 0
				for ip := 0; ip < lenPoints; ip++ {
					p := &points[ip]
					//d := dist3(i, j, k, p.x, p.y, p.z)
					var d int
					dist3_test(i, j, k, p.x, p.y, p.z, &d)
					if d < p.r {
						inRange++
					}
				}
				count++
				if inRange > inRangeMax {
					//fmt.Println(i, j, k, inRange)
					inRangeMax = inRange
					targetx = i
					targety = j
					targetz = k
				}

			}
		}
	}
	fmt.Println("inRange max", inRangeMax)
	fmt.Println(targetx, targety, targetz)
	elapsed := time.Since(start)
	fmt.Println("time", elapsed)

}

// divide by 1,000
func part2_divideT() {
	/*
		from the 10,000 function we got this
		inRange max 892
		5131 4784 2514
	*/
	filebytes := readfile("input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")

	var points = make([]_point, 0)
	for i := range filelines {
		var x, y, z, r int
		fmt.Sscanf(filelines[i],
			"pos=<%d,%d,%d>, r=%d",
			&x, &y, &z, &r)
		points = append(points, _point{x, y, z, r})
	}

	F := 1000
	for i := range points {
		p := &points[i]
		p.x /= F
		p.y /= F
		p.z /= F
		p.r /= F
	}

	/*
		junk := make([]int, 0)
		for i := range points {
			junk = append(junk, points[i].z)
		}

		sort.Ints(junk)
		fmt.Println(junk)
		// x  -14450 18712
		// y  -3192 15442
		// z  -7140 12783
	*/
	V := 100
	start := time.Now()
	count := 0
	inRangeMax := 0
	targetx := 0
	targety := 0
	targetz := 0
	lenPoints := len(points)

	for i := 51310 - V; i < 51310+V; i++ {
		for j := 47840 - V; j < 47840+V; j++ {
			for k := 25140 - V; k < 25140+V; k++ {

				inRange := 0
				for ip := 0; ip < lenPoints; ip++ {
					p := &points[ip]
					//d := dist3(i, j, k, p.x, p.y, p.z)
					var d int
					dist3_test(i, j, k, p.x, p.y, p.z, &d)
					if d < p.r {
						inRange++
					}
				}
				count++
				if inRange > inRangeMax {
					//fmt.Println(i, j, k, inRange)
					inRangeMax = inRange
					targetx = i
					targety = j
					targetz = k
				}

			}
		}
	}
	fmt.Println("inRange max", inRangeMax)
	fmt.Println(targetx, targety, targetz)
	elapsed := time.Since(start)
	fmt.Println("time", elapsed)

}

// divide by 1,00
func part2_divideH() {
	/*
		from the 1,000 function we got this
		inRange max 892
		51310 47836 25131
	*/
	filebytes := readfile("input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")

	var points = make([]_point, 0)
	for i := range filelines {
		var x, y, z, r int
		fmt.Sscanf(filelines[i],
			"pos=<%d,%d,%d>, r=%d",
			&x, &y, &z, &r)
		points = append(points, _point{x, y, z, r})
	}

	F := 100
	for i := range points {
		p := &points[i]
		p.x /= F
		p.y /= F
		p.z /= F
		p.r /= F
	}

	/*
		junk := make([]int, 0)
		for i := range points {
			junk = append(junk, points[i].x)
		}

		sort.Ints(junk)
		fmt.Println(junk)
		// x  -14450 18712
		// y  -3192 15442
		// z  -7140 12783
		return
	*/
	V := 100
	start := time.Now()
	count := 0
	inRangeMax := 0
	targetx := 0
	targety := 0
	targetz := 0
	lenPoints := len(points)

	for i := 513100 - V; i < 513100+V; i++ {
		for j := 478360 - V; j < 478360+V; j++ {
			for k := 251310 - V; k < 251310+V; k++ {

				inRange := 0
				for ip := 0; ip < lenPoints; ip++ {
					p := &points[ip]
					//d := dist3(i, j, k, p.x, p.y, p.z)
					var d int
					dist3_test(i, j, k, p.x, p.y, p.z, &d)
					if d < p.r {
						inRange++
					}
				}
				count++
				if inRange > inRangeMax {
					fmt.Println(i, j, k, inRange)
					inRangeMax = inRange
					targetx = i
					targety = j
					targetz = k
				}

			}
		}
	}
	fmt.Println("inRange max", inRangeMax)
	fmt.Println(targetx, targety, targetz)
	elapsed := time.Since(start)
	fmt.Println("time", elapsed)

}

// divide by 10
func part2_divideBy10() {
	filebytes := readfile("input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")

	var points = make([]_point, 0)
	for i := range filelines {
		var x, y, z, r int
		fmt.Sscanf(filelines[i],
			"pos=<%d,%d,%d>, r=%d",
			&x, &y, &z, &r)
		points = append(points, _point{x, y, z, r})
	}

	F := 10
	for i := range points {
		p := &points[i]
		p.x /= F
		p.y /= F
		p.z /= F
		p.r /= F
	}

	V := 100
	start := time.Now()
	count := 0
	inRangeMax := 0
	targetx := 0
	targety := 0
	targetz := 0
	lenPoints := len(points)

	for i := 5131010 - V; i < 5131010+V; i++ {
		for j := 4783570 - V; j < 4783570+V; j++ {
			for k := 2513030 - V; k < 2513030+V; k++ {

				inRange := 0
				for ip := 0; ip < lenPoints; ip++ {
					p := &points[ip]
					//d := dist3(i, j, k, p.x, p.y, p.z)
					var d int
					dist3_test(i, j, k, p.x, p.y, p.z, &d)
					if d < p.r {
						inRange++
					}
				}
				count++
				if inRange > inRangeMax {
					//fmt.Println(i, j, k, inRange)
					inRangeMax = inRange
					targetx = i
					targety = j
					targetz = k
				}

			}
		}
	}
	fmt.Println("inRange max", inRangeMax)
	fmt.Println(targetx, targety, targetz)
	elapsed := time.Since(start)
	fmt.Println("time", elapsed)

}

// divide by 1
func part2_divideBy1() {
	filebytes := readfile("input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")

	var points = make([]_point, 0)
	for i := range filelines {
		var x, y, z, r int
		fmt.Sscanf(filelines[i],
			"pos=<%d,%d,%d>, r=%d",
			&x, &y, &z, &r)
		points = append(points, _point{x, y, z, r})
	}

	F := 1
	for i := range points {
		p := &points[i]
		p.x /= F
		p.y /= F
		p.z /= F
		p.r /= F
	}

	V := 100
	start := time.Now()
	count := 0
	inRangeMax := 0
	targetx := 0
	targety := 0
	targetz := 0
	lenPoints := len(points)

	for i := 51310060 - V; i < 51310060+V; i++ {
		for j := 47835730 - V; j < 47835730+V; j++ {
			for k := 25130320 - V; k < 25130320+V; k++ {

				inRange := 0
				for ip := 0; ip < lenPoints; ip++ {
					p := &points[ip]
					//d := dist3(i, j, k, p.x, p.y, p.z)
					var d int
					dist3_test(i, j, k, p.x, p.y, p.z, &d)
					if d < p.r {
						inRange++
					}
				}
				count++
				if inRange > inRangeMax {
					fmt.Println(i, j, k, inRange)
					inRangeMax = inRange
					targetx = i
					targety = j
					targetz = k
				}

			}
		}
	}
	fmt.Println("inRange max", inRangeMax)
	fmt.Println(targetx, targety, targetz)
	elapsed := time.Since(start)
	fmt.Println("time", elapsed)

}

func part2_final() {
	// guessed this and it was too high
	//124 276 104

	// guess this and it was too low
	//27 608 415

	//fmt.Println(dist3(0, 0, 0, 51310045, 47835736, 25130323))

	filebytes := readfile("input")
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")

	var points = make([]_point, 0)
	for i := range filelines {
		var x, y, z, r int
		fmt.Sscanf(filelines[i],
			"pos=<%d,%d,%d>, r=%d",
			&x, &y, &z, &r)
		points = append(points, _point{x, y, z, r})
	}
	count := 0
	shortest := math.MaxInt32
	for i := range points {
		p := points[i]
		d := dist3(51310045, 47835736, 25130323, p.x, p.y, p.z)
		if d <= p.r {
			fmt.Println("need to check this one")

			dist_to_zero := dist3(0, 0, 0, p.x, p.y, p.z)
			if dist_to_zero < shortest {
				shortest = dist_to_zero
			}
			count++
		}
	}
	fmt.Println("shortest", shortest)
}

func main() {
	//part1()
	//part2_divideM()
	//part2_divideHT()
	//part2_divideTT()
	//part2_divideT()
	//part2_divideH()
	//part2_divideBy10()
	//part2_divideBy1()
	part2_final()
}
