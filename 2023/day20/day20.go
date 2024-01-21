package aoc2023

import (
	"advent-of-code/util/math"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

type Module interface {
	send(event PulseEvent) []PulseEvent
	getPeers() []string
	getType() string
}

type FlipFlop struct {
	name  string
	peers []string
	state bool
}

type Conjunction struct {
	name   string
	peers  []string
	states map[string]bool
}

type Broadcaster struct {
	name  string
	peers []string
}

func (b *Broadcaster) send(event PulseEvent) []PulseEvent {
	result := []PulseEvent{}
	for _, peer := range b.peers {
		result = append(result,
			PulseEvent{source: b.name, dst: peer, pulse: event.pulse})
	}
	return result
}

func (b *Broadcaster) getPeers() []string {
	return b.peers
}

func (b *Broadcaster) getType() string {
	return "Broadcaster"
}

func (c *Conjunction) send(event PulseEvent) []PulseEvent {
	c.states[event.source] = event.pulse
	states := maps.Values(c.states)
	hasLow := slices.Contains(states, false)
	result := []PulseEvent{}
	for _, peer := range c.peers {
		result = append(result,
			PulseEvent{source: c.name, dst: peer, pulse: hasLow})
	}
	return result
}

func (c *Conjunction) getPeers() []string {
	return c.peers
}

func (c *Conjunction) getType() string {
	return "Conjunction"
}

func (f *FlipFlop) send(event PulseEvent) []PulseEvent {
	result := []PulseEvent{}
	if !event.pulse {
		f.state = !f.state
		for _, peer := range f.peers {
			result = append(result,
				PulseEvent{source: f.name, dst: peer, pulse: f.state})
		}
	}
	return result
}

func (f *FlipFlop) getPeers() []string {
	return f.peers
}

func (f *FlipFlop) getType() string {
	return "FlipFlop"
}

type PulseEvent struct {
	source string
	dst    string
	pulse  bool
}

func CountButtonPushes(input []string) int {
	modules := ParseData(input)
	rxRoot := FindRootPuleses(modules, "rx")[0]
	rootsOfRoot := FindRootPuleses(modules, rxRoot)
	occurances := map[string]int{}

	queue := []PulseEvent{}
	counter := 0
	for {
		counter++

		start := modules["broadcaster"]
		startEvents := start.send(PulseEvent{pulse: false})
		queue = append(queue, startEvents...)

		for len(queue) > 0 {
			currEvent := queue[0]
			queue = queue[1:]
			currModule := modules[currEvent.dst]
			if slices.Contains(rootsOfRoot, currEvent.dst) && !currEvent.pulse {
				occurances[currEvent.dst] = counter
				if len(occurances) == len(rootsOfRoot) {
					return CalculateResult(occurances)
				}
			}
			if currModule == nil {
				continue
			}
			newEvents := currModule.send(currEvent)
			queue = append(queue, newEvents...)
		}
	}
}

func CalculateResult(occurances map[string]int) int {
	return math.LCM(maps.Values(occurances)...)
}

func FindRootPuleses(modules map[string]Module, names ...string) []string {
	result := []string{}
	for name, module := range modules {
		for _, peer := range module.getPeers() {
			if slices.Contains(names, peer) {
				result = append(result, name)
			}
		}
	}
	return result
}

func PushButton(input []string) int {
	times := 1000
	modules := ParseData(input)

	lowCount, highCount := times, 0
	queue := []PulseEvent{}

	for i := 0; i < times; i++ {
		start := modules["broadcaster"]
		startEvents := start.send(PulseEvent{pulse: false})
		queue = append(queue, startEvents...)

		for len(queue) > 0 {
			currEvent := queue[0]
			if currEvent.pulse {
				highCount++
			} else {
				lowCount++
			}
			queue = queue[1:]
			currModule := modules[currEvent.dst]
			if currModule == nil {
				continue
			}
			newEvents := currModule.send(currEvent)
			queue = append(queue, newEvents...)
		}
	}
	return lowCount * highCount
}

func ParseData(input []string) map[string]Module {
	modules := map[string]Module{}
	for _, line := range input {
		data := strings.Split(line, " -> ")
		moduleNames := strings.Split(data[1], ", ")
		name := data[0]
		if data[0][0] == '%' {
			modules[name[1:]] = &FlipFlop{name: name[1:], peers: moduleNames, state: false}
		} else if data[0][0] == '&' {
			modules[name[1:]] = &Conjunction{name: name[1:], peers: moduleNames, states: map[string]bool{}}
		} else if name == "broadcaster" {
			modules[name] = &Broadcaster{name: name, peers: moduleNames}
		}
	}

	for name, module := range modules {
		for _, peer := range module.getPeers() {
			var dest interface{} = modules[peer]
			c, ok := dest.(*Conjunction)
			if ok {
				c.states[name] = false
				modules[peer] = &Conjunction{
					name: peer, peers: c.peers, states: c.states}
			}
		}
	}

	return modules
}
