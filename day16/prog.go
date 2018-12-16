package main

import (
	"fmt"
	"io/ioutil"
)

type _instr struct {
	opcode int
	a      int
	b      int
	c      int
}
type _registers [4]int

func readfile(filename string) []byte {
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("could not read file")
		return nil
	}
	return filedata
}

func addr(inreg _registers, oper _instr) {

}
func addi(inreg _registers, oper _instr) {

}

func test_addi() {
	inreg := _registers{3, 2, 1, 1}
	addi(inreg, _instr{9, 2, 1, 2})
	//? profit

}

func main() {
}
