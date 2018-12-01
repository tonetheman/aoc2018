package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const size = 64

type bits uint64
type BitSet []bits

func (s *BitSet) Set(i uint) {
	if len(*s) < int(i/size+1) {
		r := make([]bits, i/size+1)
		copy(r, *s)
		*s = r
	}
	(*s)[i/size] |= 1 << (i % size)
}
func (s *BitSet) Clear(i uint) {
	if len(*s) >= int(i/size+1) {
		(*s)[i/size] &^= 1 << (i % size)
	}
}
func (s *BitSet) IsSet(i uint) bool {
	return (*s)[i/size]&(1<<(i%size)) != 0
}

// random idea about using a bloom filter
func h(a int64) int32 {
	a = (a + 0x7ed55d16) + (a << 12)
	a = (a ^ 0xc761c23c) ^ (a >> 19)
	a = (a + 0x165667b1) + (a << 5)
	a = (a + 0xd3a2646c) ^ (a << 9)
	a = (a + 0xfd7046c5) + (a << 3)
	a = (a ^ 0xb55a4f09) ^ (a >> 16)
	return int32(a)
}

func stringsToInts(s []string) []int {
	var res []int
	for _, t := range s {
		val, _ := strconv.Atoi(t)
		res = append(res, val)
	}
	return res
}

func part1(a []int) int {
	freq := 0
	for _, val := range a {
		freq += val
	}
	return freq
}

func part2(a []int) int {
	index := 0
	freq := 0
	m := make(map[int]bool)
	m[freq] = true
	for {
		val := a[index]
		freq += val
		// check if the value is there
		_, ok := m[freq]
		if ok {
			fmt.Println("dup", freq)
			break
		} else {
			m[freq] = true
		}
		index++
		if index >= len(a) {
			index = 0
		}
	}
	return 0
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(h(int64(i)) % 63)
	}

	s := new(BitSet)
	fmt.Println(s)

	filedata, err := ioutil.ReadFile("..\\data_day1.txt")
	if err != nil {
		fmt.Println("could not read file")
		return
	}

	// this might be a DOS thing the \r\n
	data := strings.Split(string(filedata), "\r\n")
	fmt.Println(data)
	idata := stringsToInts(data)
	fmt.Println(idata)
	fmt.Println(part1(idata))
	fmt.Println(part2(idata))
}
