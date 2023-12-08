package aoc2017

import (
	"strconv"
	"strings"
	"sync"
	"time"
)

func Start(input []string) int {
	channel0 := make(chan int, 1000)
	channel1 := make(chan int, 1000)

	program0 := Program{name: 0, registers: map[string]int{}, instructions: input, sendChannel: channel0, receiveChannel: channel1}
	program1 := Program{name: 1, registers: map[string]int{}, instructions: input, sendChannel: channel1, receiveChannel: channel0}

	var wg sync.WaitGroup
	wg.Add(2)
	go program0.Run(&wg)
	go program1.Run(&wg)

	wg.Wait()
	return program1.sendCount
}

type Program struct {
	name           int
	registers      map[string]int
	instructions   []string
	sendChannel    chan int
	receiveChannel chan int
	sendCount      int
}

func (p *Program) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	p.registers["p"] = p.name

	index := 0
	for index < len(p.instructions) {
		line := p.instructions[index]
		data := strings.Split(line, " ")
		switch data[0] {
		case "snd":
			value := getValue(data[1], p.registers)
			p.sendChannel <- value
			p.sendCount++
		case "set":
			value := getValue(data[2], p.registers)
			p.registers[data[1]] = value
		case "add":
			value := getValue(data[2], p.registers)
			p.registers[data[1]] += value
		case "mul":
			value := getValue(data[2], p.registers)
			p.registers[data[1]] *= value
		case "mod":
			value := getValue(data[2], p.registers)
			if value != 0 {
				p.registers[data[1]] %= value
			}
		case "rcv":
			select {
			case value := <-p.receiveChannel:
				p.registers[data[1]] = value
			case <-time.After(2 * time.Second):
				return
			}
		case "jgz":
			jmpCond := getValue(data[1], p.registers)
			jmpSize := getValue(data[2], p.registers)

			if jmpCond > 0 {
				index += jmpSize - 1
			}
		}
		index++
	}
}

func RecoverLastSound(input []string) int {
	registers := map[string]int{}

	sound := 0
	for index := 0; index < len(input); index++ {
		line := input[index]
		data := strings.Split(line, " ")
		cmd := data[0]
		switch cmd {
		case "snd":
			sound = registers[data[1]]
		case "set":
			value := getValue(data[2], registers)
			registers[data[1]] = value

		case "add":
			value := getValue(data[2], registers)
			registers[data[1]] += value

		case "mul":
			value := getValue(data[2], registers)
			registers[data[1]] *= value
		case "mod":
			value := getValue(data[2], registers)
			if value != 0 {
				registers[data[1]] %= value
			} else {
				panic("Modulo of 0")
			}
		case "rcv":
			if registers[data[1]] > 0 {
				return sound
			}
		case "jgz":
			jmpCond := getValue(data[1], registers)
			jmpSize := getValue(data[2], registers)

			if jmpCond > 0 {
				index += jmpSize - 1
			}
		}
	}
	panic("Something wrong!")
}

func getValue(s string, registers map[string]int) int {
	if value, err := strconv.Atoi(s); err == nil {
		return value
	}
	return registers[s]
}
