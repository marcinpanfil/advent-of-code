package aoc2017

import (
	"advent-of-code/util/array"
	"fmt"
	"strconv"
	"strings"
)

func Dance(programs []string, moves []string, times int) string {
	for i := 0; i < times; i++ {
		for _, move := range moves {
			action := string(move[0])
			if action == "s" {
				idx, _ := strconv.Atoi(move[1:])
				tmp := programs[len(programs)-idx:]
				tmp = append(tmp, programs[0:len(programs)-idx]...)
				programs = tmp
			} else if action == "x" {
				posStr := strings.Split(move[1:], "/")
				posA, _ := strconv.Atoi(posStr[0])
				posB, _ := strconv.Atoi(posStr[1])
				tmp := programs[posA]
				programs[posA] = programs[posB]
				programs[posB] = tmp
			} else if action == "p" {
				posStr := strings.Split(move[1:], "/")
				posA := array.IndexOf[string](programs, posStr[0])
				posB := array.IndexOf[string](programs, posStr[1])
				tmp := programs[posA]
				programs[posA] = programs[posB]
				programs[posB] = tmp
			} else {
				panic("Wrong action")
			}
		}
		fmt.Printf("%v %s\n", i, strings.Join(programs, ""))
	}
	return strings.Join(programs, "")
}
