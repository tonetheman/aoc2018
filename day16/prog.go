package main

import (
	"fmt"
	"io/ioutil"
)

type _instr struct {
	opcode int
	a      int
	b      int
	output int
}
type _registers [4]int

func (r *_registers) copy(src _registers) {
	r[0] = src[0]
	r[1] = src[1]
	r[2] = src[2]
	r[3] = src[3]
}

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
func addi(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	res[oper.output] = inreg[oper.a] + oper.b
	return res
}

func mulr(inreg _registers, oper _instr) {
	inreg[oper.output] = inreg[oper.a] * inreg[oper.b]
}

func seti(inreg _registers, oper _instr) {
	inreg[oper.output] = inreg[oper.a]
}

func testAddi() {
	fmt.Println("testAddi")
	inreg := _registers{3, 2, 1, 1}
	fmt.Println("addi enter", inreg)
	addi(inreg, _instr{9, 2, 1, 2})
	fmt.Println("addi exit", inreg)
}
func testMulr() {
	fmt.Println("testMulr")
	inreg := _registers{3, 2, 1, 1}
	fmt.Println("mulr enter", inreg)
	mulr(inreg, _instr{9, 2, 1, 2})
	fmt.Println("mulr exit", inreg)
}
func testSeti() {
	fmt.Println("testSeti")
	inreg := _registers{3, 2, 1, 1}
	fmt.Println("seti enter", inreg)
	seti(inreg, _instr{9, 2, 1, 2})
	fmt.Println("seti exit", inreg)
}

func examples() {
	testAddi()
	testMulr()
	testSeti()
}

func main() {
}
