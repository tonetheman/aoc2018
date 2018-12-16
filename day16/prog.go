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

func addr(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	res[oper.output] = inreg[oper.a] + inreg[oper.b]
	return res
}
func addi(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	res[oper.output] = inreg[oper.a] + oper.b
	return res
}
func mulr(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	res[oper.output] = inreg[oper.a] * inreg[oper.b]
	return res
}
func muli(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	res[oper.output] = inreg[oper.a] * oper.b
	return res
}

// bitwise and
func banr(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	res[oper.output] = inreg[oper.a] & inreg[oper.b]
	return res
}
func bani(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	res[oper.output] = inreg[oper.a] & oper.b
	return res
}

// bitwise or
func borr(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	res[oper.output] = inreg[oper.a] | inreg[oper.b]
	return res
}
func bori(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	res[oper.output] = inreg[oper.a] | oper.b
	return res
}

// assignment
func seti(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	res[oper.output] = oper.a // value
	return res
}
func setr(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	res[oper.output] = inreg[oper.a]
	return res
}

// greater than testing
func gtir(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	if oper.a > inreg[oper.b] {
		res[oper.output] = 1
	} else {
		res[oper.output] = 0
	}
	return res
}
func gtri(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	if inreg[oper.a] > oper.b {
		res[oper.output] = 1
	} else {
		res[oper.output] = 0
	}
	return res
}
func gtrr(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	if inreg[oper.a] > inreg[oper.b] {
		res[oper.output] = 1
	} else {
		res[oper.output] = 0
	}
	return res
}

// equality testing
func eqir(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	if oper.a == inreg[oper.b] {
		res[oper.output] = 1
	} else {
		res[oper.output] = 0
	}
	return res
}
func eqri(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	if inreg[oper.a] == oper.b {
		res[oper.output] = 1
	} else {
		res[oper.output] = 0
	}
	return res
}
func eqrr(inreg _registers, oper _instr) _registers {
	res := _registers{}
	res.copy(inreg)
	if inreg[oper.a] == inreg[oper.b] {
		res[oper.output] = 1
	} else {
		res[oper.output] = 0
	}
	return res
}

func testAddi() {
	fmt.Println("testAddi")
	inreg := _registers{3, 2, 1, 1}
	fmt.Println("addi enter", inreg)
	outreg := addi(inreg, _instr{9, 2, 1, 2})
	fmt.Println("addi exit", outreg)
}
func testMulr() {
	fmt.Println("testMulr")
	inreg := _registers{3, 2, 1, 1}
	fmt.Println("mulr enter", inreg)
	outreg := mulr(inreg, _instr{9, 2, 1, 2})
	fmt.Println("mulr exit", outreg)
}
func testSeti() {
	fmt.Println("testSeti")
	inreg := _registers{3, 2, 1, 1}
	fmt.Println("seti enter", inreg)
	outreg := seti(inreg, _instr{9, 2, 1, 2})
	fmt.Println("seti exit", outreg)
}

func examples() {
	testAddi()
	testMulr()
	testSeti()
}

func main() {
	examples()
}
