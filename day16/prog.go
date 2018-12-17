package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type _instr struct {
	opcode int
	a      int
	b      int
	output int
}
type _registers [4]int
type ioper func(inreg _registers, oper _instr) _registers

func (r *_registers) copy(src _registers) {
	r[0] = src[0]
	r[1] = src[1]
	r[2] = src[2]
	r[3] = src[3]
}
func (r *_registers) verify(v _registers) bool {
	return (v[0] == r[0]) &&
		(v[1] == r[1]) &&
		(v[2] == r[2]) &&
		(v[3] == r[3])
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

func iTesting() {
	names := []string{"addr", "addi", "mulr", "muli",
		"banr", "bani", "borr", "bori",
		"seti", "setr", "gtir", "gtri", "gtrr", "eqir", "eqri", "eqrr"}
	todos := []ioper{addr, addi, mulr, muli, banr, bani, borr, bori,
		seti, setr, gtir, gtri, gtrr, eqir, eqri, eqrr}
	for i := range todos {
		inreg := _registers{3, 2, 1, 1}
		outreg := todos[i](inreg, _instr{9, 2, 1, 2})
		target := _registers{3, 2, 2, 1}
		if target.verify(outreg) {
			fmt.Println(names[i], outreg)
		}

	}

}

func check(inreg _registers, outreg _registers, op _instr) map[string]_registers {
	m := make(map[string]_registers)

	names := []string{"addr", "addi", "mulr", "muli",
		"banr", "bani", "borr", "bori",
		"seti", "setr", "gtir", "gtri", "gtrr", "eqir", "eqri", "eqrr"}
	todos := []ioper{addr, addi, mulr, muli, banr, bani, borr, bori,
		seti, setr, gtir, gtri, gtrr, eqir, eqri, eqrr}
	for i := range todos {
		// make a copy of inreg
		m[names[i]] = todos[i](inreg, op)
	}
	return m
}

func part1() {
	filebytes := readfile("input-samples")
	for i := range filebytes {
		// who knows if this works
		if filebytes[i] == '\r' {
			filebytes[i] = ' '
		}
	}
	filestring := string(filebytes)
	lines := strings.Split(filestring, "\n")
	for i := range lines {
		lines[i] = strings.Trim(lines[i], " \n")
	}
	var rr _registers
	var aa _registers
	var ii _instr
	for i := range lines {
		if strings.HasPrefix(lines[i], "Before:") {
			fmt.Sscanf(lines[i],
				"Before: [%d, %d, %d, %d]", &rr[0], &rr[1], &rr[2], &rr[3])
		} else if strings.HasPrefix(lines[i], "After:") {
			fmt.Sscanf(lines[i],
				"After:  [%d, %d, %d, %d]", &aa[0], &aa[1], &aa[2], &aa[3])

			// DO WORK HERE

			check(rr, aa, ii)

		} else {
			if len(lines[i]) == 0 {
				// ignore this screw EOL in windows
			} else {
				//fmt.Println("WUT?", lines[i])
				fmt.Sscanf(lines[i],
					"%d %d %d %d", &ii.opcode, &ii.a, &ii.b, &ii.output)
			}
		}
	}
}

func a() {

	inreg := _registers{3, 2, 1, 1}
	outreg := _registers{}
	m := check(inreg, outreg, _instr{9, 2, 1, 2})
	res := make(map[_registers]int)
	for k, v := range m {
		fmt.Println(k, v)
		res[v]++
	}
	fmt.Println(res)
}
func b() {
	inreg := _registers{0, 1, 2, 1}
	outreg := _registers{}
	m := check(inreg, outreg, _instr{14, 1, 3, 3})
	res := make(map[_registers]int)
	for k, v := range m {
		fmt.Println(k, v)
		res[v]++
	}
	fmt.Println(res)
}
func main() {
	b()
}
