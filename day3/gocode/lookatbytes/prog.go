package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

func readfile(filename string) []byte {
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("could not read file")
		return nil
	}
	return filedata
}

func bytestostring(buffer []byte) string {
	return string(buffer)
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Hello from Windows")
	}

	buffer := readfile("..\\..\\input")
	fmt.Println(buffer[0:32])
	for i := 0; i < len(buffer); i++ {
		if buffer[i] == '\r' {
			buffer[i] = ' '
		}
	}
	s := bytestostring(buffer)
	for i := 0; i < 32; i++ {
		fmt.Printf("%x, ", s[i])
	}
	fmt.Println()
}
