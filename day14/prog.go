package main

import (
	"fmt"
	"strconv"
)

const BS = 919901 + 200

type _board struct {
	elf1 int
	elf2 int
	pos  int
	data [BS]int
}

func addDigits(val int, b *_board) {
	s := strconv.Itoa(val)
	for _, v := range s {
		res, _ := strconv.Atoi(string(v))
		b.data[b.pos] = res
		b.pos++
	}
}

func (b _board) print() {
	fmt.Println("elf pos", b.elf1, b.elf2)
	for i := 0; i < b.pos; i++ {
		if i == b.elf1 {
			fmt.Print("1")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	for i := 0; i < b.pos; i++ {
		if i == b.elf2 {
			fmt.Print("2")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	for i := 0; i < b.pos; i++ {
		fmt.Print(b.data[i])
	}
	fmt.Println()
}

func (b *_board) newrecipe() int {
	return b.data[b.elf1] + b.data[b.elf2]
}

func part1() {
	var board _board
	addDigits(37, &board)

	board.elf1 = 0 // index
	board.elf2 = 1 // index
	input := 919901
	count := 0
	for {
		newrecipe := board.newrecipe() //board.data[elf1] + board.data[elf2]
		//fmt.Println("newrecipe is", newrecipe)
		addDigits(newrecipe, &board)
		elf1_moves := board.data[board.elf1] + 1
		elf2_moves := board.data[board.elf2] + 1
		board.elf1 = (board.elf1 + elf1_moves) % board.pos
		board.elf2 = (board.elf2 + elf2_moves) % board.pos
		//fmt.Println("------------------")
		//fmt.Println("count", count)
		//board.print()
		count++
		//if count == 900 {
		//	break
		//}
		if board.pos-9 > input {
			fmt.Println("cause done")
			break
		}
	}
	for i := input; i <= input+9; i++ {
		fmt.Print(board.data[i])
	}
	fmt.Println()

}

func main() {
	example()
}
