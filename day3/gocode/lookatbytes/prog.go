package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
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
	sdata := strings.Split(s, "\n")
	fmt.Println(sdata)
	fmt.Println("FIRST", sdata[0])

	for i := 0; i < len(sdata); i++ {
		var id, posx, posy, w, h int
		fmt.Sscanf(sdata[i], "#%d @ %d,%d: %dx%d", &id, &posx, &posy, &w, &h)
		fmt.Println(id, posx, posy, w, h)
	}
	fmt.Println("fin")
}
