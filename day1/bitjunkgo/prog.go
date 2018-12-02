package main

import "fmt"

type Bits []uint64

var bits Bits

func (b *Bits) clear() {
	fmt.Println(b)
}
func (b *Bits) set(i int) {
	(*b)[0] = uint64(i)
}

func main() {
	bits = make([]uint64, 1)

	bits.set(60)
	fmt.Println(bits)

	bits.clear()
	fmt.Println(bits)

}
