package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type cloth struct {
	id            int
	posx, posy    int
	width, height int
}

func readfile(filename string) []byte {
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("could not read file")
		return nil
	}
	return filedata
}

func bytesToString(buffer []byte) string {
	return string(buffer)
}

func simpleparseall(sdata []string) []cloth {
	res := make([]cloth, 0)
	for i := 0; i < len(sdata); i++ {
		var id, posx, posy, w, h int
		fmt.Sscanf(sdata[i], "#%d @ %d,%d: %dx%d", &id, &posx, &posy, &w, &h)
		//fmt.Println(id, posx, posy, w, h)
		res = append(res, cloth{id, posx, posy, w, h})
	}
	return res
}

func main() {
	inputdata := readfile("..\\..\\input")
	stringdata := bytesToString(inputdata)
	splitdata := strings.Split(stringdata, "\n")
	cloths := simpleparseall(splitdata)
	fmt.Printf("{ \"data\" : \n")
	for i := 0; i < len(cloths); i++ {
		c := cloths[i]
		fmt.Printf("[%d,%d,%d,%d,%d],\n", c.id, c.posx, c.posy, c.width, c.height)
	}
	fmt.Printf("}\n")
}
