package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type _rec struct {
	y, m, d, hh, mm int
	desc            string
}

type byTimestamp []_rec

func (a byTimestamp) Len() int {
	return len(a)
}
func (a byTimestamp) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a byTimestamp) Less(i, j int) bool {
	if a[i].y != a[j].y {
		return a[i].y < a[j].y
	}
	if a[i].m != a[j].m {
		return a[i].m < a[j].m
	}
	if a[i].d != a[j].d {
		return a[i].d < a[j].d
	}
	if a[i].hh != a[j].hh {
		return a[i].hh < a[j].hh
	}
	return a[i].mm < a[j].mm
}

func readfile(filename string) []byte {
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("could not read file")
		return nil
	}
	return filedata
}

func getInputRecords(filename string) []_rec {
	filebytes := readfile(filename)
	// fix windows/dos
	for i := 0; i < len(filebytes); i++ {
		if filebytes[i] == '\r' {
			filebytes[i] = ' '
		}
	}
	filestring := string(filebytes)
	filelines := strings.Split(filestring, "\n")

	_recs := make([]_rec, 0)
	for i := 0; i < len(filelines); i++ {
		var r _rec
		// this will get timestamp on front
		fmt.Sscanf(filelines[i], "[%d-%d-%d %d:%d]",
			&r.y, &r.m, &r.d, &r.hh, &r.mm)

		_tmplinedata := strings.Split(filelines[i], "]")
		r.desc = strings.Trim(_tmplinedata[1], " ")
		//fmt.Println(filelines[i])
		//fmt.Println(r)
		_recs = append(_recs, _rec{r.y, r.m, r.d, r.hh, r.mm, r.desc})
	}
	return _recs
}

func pr(recs []_rec) {
	for i := 0; i < len(recs); i++ {
		fmt.Println(i, recs[i])
	}
}

func pr10(recs []_rec) {
	for i := 0; i < 10; i++ {
		fmt.Println(recs[i])
	}
}

func validate_hour(recs []_rec) {
	var m map[int]int = make(map[int]int)
	for i := 0; i < len(recs); i++ {
		r := recs[i]
		m[r.hh]++
	}
	fmt.Println(m)
	// PROVED: hours only 23 and 0 in the file for real input
	// and example input
}

// checks for all zero
func allzero(a [60]int) bool {
	res := true
	for i := 0; i < 60; i++ {
		if a[i] != 0 {
			res = false
		}
	}
	return res
}

// AKA part1
func whosleptthemost(recs []_rec) {
	current_guard_id := 0
	sleep_start_hour := 0
	sleep_start_min := 0

	var guard_sleep_totals map[int]int = make(map[int]int)
	for i := 0; i < len(recs); i++ {
		r := recs[i]
		if strings.HasPrefix(r.desc, "Guard #") {
			// begin of shift record
			fmt.Sscanf(r.desc, "Guard #%d", &current_guard_id)
			fmt.Println("current guard id changed to", current_guard_id)
		}
		if strings.HasPrefix(r.desc, "falls") {
			sleep_start_hour = r.hh
			sleep_start_min = r.mm
		}
		if strings.HasPrefix(r.desc, "wakes") {
			// compute sleep now
			//fmt.Println("guard", current_guard_id, "slept", (r.mm - sleep_start_min))
			_, ok := guard_sleep_totals[current_guard_id]
			if !ok {
				guard_sleep_totals[current_guard_id] = (r.mm - sleep_start_min)
			} else {
				guard_sleep_totals[current_guard_id] += (r.mm - sleep_start_min)
			}
			if sleep_start_hour != r.hh {
				fmt.Println("WARN")
			}
		}
	}

	whototal := -1
	who := -1
	for k, v := range guard_sleep_totals {
		if v > whototal {
			who = k
			whototal = v
		}
	}
	fmt.Println("who slept the most", who, whototal)
	var tmp [60]int
	var total_tmp [60]int

	// now find the most slept min for the dude
	fmt.Println("now looking for when this dude slept...")
	for i := 0; i < len(recs); i++ {
		// get a pointer to the current record
		r := recs[i]
		if strings.HasPrefix(r.desc, "Guard #") {
			// begin of shift record
			fmt.Sscanf(r.desc, "Guard #%d", &current_guard_id)
			if current_guard_id == who {
				if allzero(tmp) {

				} else {
					fmt.Println("RES", tmp)
					for k := 0; k < 60; k++ {
						total_tmp[k] += tmp[k]
					}
				}

				fmt.Println("current guard id changed to", current_guard_id, r)
				for j := 0; j < 60; j++ {
					tmp[j] = 0
				}
			}
		}
		if strings.HasPrefix(r.desc, "falls") {
			if current_guard_id == who {
				// this is the guy we are intereted in
				fmt.Println("\t", r)
				sleep_start_hour = r.hh
				sleep_start_min = r.mm
			}
		}
		if strings.HasPrefix(r.desc, "wakes") {
			if current_guard_id == who {
				// this is the guy we are interested in
				fmt.Println("\t", r)
				for j := sleep_start_min; j < r.mm; j++ {
					tmp[j] = 1
				}
			}
		}
	}
	fmt.Println("RES", tmp)
	for k := 0; k < 60; k++ {
		total_tmp[k] += tmp[k]
	}
	fmt.Println("TOTALRES", total_tmp)
	largest_index := 0
	largest_value := -1
	for i := 0; i < 60; i++ {
		if total_tmp[i] > largest_value {
			largest_index = i
			largest_value = total_tmp[i]
		}
	}
	fmt.Println("largest value and index", largest_value, largest_index)
	fmt.Println("who, largestindex and final answer",
		who, largest_index, who*largest_index)
}

func part2(recs []_rec) {
	var current_guard_id int
	var last_guard_id int = -1
	var sleep_start_hour, sleep_start_min int
	var tmp [60]int

	cleartmp := func() {
		for i := 0; i < 60; i++ {
			tmp[i] = 0
		}
	}
	cleartmp()

	type guard_and_seconds struct {
		id  int
		tmp [60]int
	}

	var alltimes []guard_and_seconds = make([]guard_and_seconds, 0)
	makeone := func() guard_and_seconds {
		var crud guard_and_seconds
		crud.id = last_guard_id
		for i := 0; i < 60; i++ {
			crud.tmp[i] = tmp[i]
		}
		return crud
	}
	for i := 0; i < len(recs); i++ {
		// get a pointer to the current record
		r := recs[i]
		if strings.HasPrefix(r.desc, "Guard #") {
			// begin of shift record
			fmt.Sscanf(r.desc, "Guard #%d", &current_guard_id)
			//fmt.Println("current guard id changed to", current_guard_id, r)

			if last_guard_id == -1 {
				// this is the first guard
				last_guard_id = current_guard_id
			} else {
				// there was a last guard id
				//+fmt.Println(last_guard_id, tmp)
				alltimes = append(alltimes, makeone())
				last_guard_id = current_guard_id
				cleartmp()
			}

		}
		if strings.HasPrefix(r.desc, "falls") {
			//fmt.Println("\t", r)
			sleep_start_hour = r.hh
			sleep_start_min = r.mm
		}
		if strings.HasPrefix(r.desc, "wakes") {
			//fmt.Println("\t", r)
			for j := sleep_start_min; j < r.mm; j++ {
				tmp[j] += 1
			}
		}
	}

	//+ fmt.Println(last_guard_id, tmp)
	alltimes = append(alltimes, makeone())

	// GO BUG!!!!!!!! had to put this here
	fmt.Println(sleep_start_hour, sleep_start_min)

	for i := 0; i < len(alltimes); i++ {
		fmt.Println(alltimes[i])
	}

	// need to gather times now according to guard
	var gatherd map[int][60]int = make(map[int][60]int, 0)
	for i := 0; i < len(alltimes); i++ {
		a := alltimes[i]
		g := gatherd[a.id]
		for j := 0; j < 60; j++ {
			g[j] += a.tmp[j]
		}
		gatherd[a.id] = g
	}
	fmt.Println("--------------")
	sleepiest_guard := -1
	sleepiest_min := -1
	sleepiest_index := -1
	for k, v := range gatherd {
		fmt.Println(k, v)
		for index := 0; index < 60; index++ {
			if v[index] > sleepiest_min {
				sleepiest_guard = k
				sleepiest_min = v[index]
				sleepiest_index = index
				//fmt.Println("INDEX", index)
			}
		}
	}
	fmt.Println("sleepiest guard and min", sleepiest_guard, sleepiest_min)
	fmt.Println("sleepiest index", sleepiest_index)
	fmt.Println("res", sleepiest_guard*sleepiest_index)
}

func main() {
	recs := getInputRecords("../input")
	//recs := getInputRecords("./input-example")
	junk := byTimestamp(recs)
	sort.Sort(junk)
	//pr(recs)
	//whosleptthemost(recs)
	part2(recs)

}
