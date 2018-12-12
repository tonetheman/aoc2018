package main

import "fmt"

func computeplvl(x int, y int, gridserial int) int {
	rackid := x + 10
	startingpowerlevel := rackid * y
	startingpowerlevel += gridserial
	startingpowerlevel *= rackid
	hdigit := int(startingpowerlevel/100) % 10
	return hdigit - 5
}

func test_plvl() {

	fmt.Println(computeplvl(3, 5, 8))
	fmt.Println(computeplvl(122, 79, 57))
	fmt.Println(computeplvl(217, 196, 39))
	fmt.Println(computeplvl(101, 153, 71))
}

func grid(gridserial int) [300][300]int {
	var grid [300][300]int

	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			grid[i][j] = computeplvl(i+1, j+1, gridserial)
		}
	}
	return grid
}

func findBest3x3(grid [300][300]int) (int, int, int) {
	maxtmp := 0
	x := -1
	y := -1
	for i := 0; i < 300-3; i++ {
		for j := 0; j < 300-3; j++ {
			tmp := grid[i][j] + grid[i][j+1] + grid[i][j+2]
			tmp += grid[i+1][j] + grid[i+1][j+1] + grid[i+1][j+2]
			tmp += grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]
			if tmp > maxtmp {
				maxtmp = tmp
				x = i + 1
				y = j + 1
			}
		}
	}
	return x, y, maxtmp
}

func testgrid1() {
	g := grid(18)
	fmt.Println(findBest3x3(g))
}

func testgrid2() {
	g := grid(42)
	fmt.Println(findBest3x3(g))
}

func part1() {
	g := grid(2187)
	fmt.Println(findBest3x3(g))
}

func main() {
	part1()
}
