package aoc2023

import (
	"advent-of-code/util/array"
	"fmt"
	"sort"
	"strings"

	"golang.org/x/exp/maps"
)

type MapEntry struct {
	Key   string
	Value int
}

func Find6WiresToRemove(input []string) int {
	connections := ParseInput(input)

	occurrences := map[string]int{}
	wires := maps.Keys(connections)
	for i := 0; i < 15; i++ {
		start := wires[i]
		for j := 0; j < len(wires); j++ {
			end := wires[j]
			path := FindPath(connections, start, end)
			for _, v := range path {
				count := occurrences[v]
				occurrences[v] = count + 1
			}
		}
	}

	mostVisited := GetMostVisited(occurrences)
	for _, k := range mostVisited {
		delete(connections, k.Key)
	}

	graphSizes := CalculateGraphSizes(connections)
	result := 1
	for _, size := range graphSizes {
		result *= size
	}
	return result
}

func GetMostVisited(occurrences map[string]int) []MapEntry {
	var sortedOccurrences []MapEntry
	for k, v := range occurrences {
		sortedOccurrences = append(sortedOccurrences, MapEntry{k, v})
	}

	sort.Slice(sortedOccurrences, func(i, j int) bool {
		return sortedOccurrences[i].Value > sortedOccurrences[j].Value
	})
	return sortedOccurrences[:6]
}

func CalculateGraphSizes(connections map[string][]string) map[string]int {
	visited := make(map[string]bool)
	graphSizes := make(map[string]int)
	count := 0

	for wire := range connections {
		if !visited[wire] {
			count++
			size := TraverseGraph(wire, connections, visited)
			graphSizes[fmt.Sprintf("Graph %d", count)] = size
		}
	}
	return graphSizes
}

func TraverseGraph(wire string, connections map[string][]string, visited map[string]bool) int {
	size := 0
	queue := []string{wire}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if !visited[curr] {
			visited[curr] = true
			size++

			neighbors := connections[curr]
			for _, neighbor := range neighbors {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return size
}

func FindPath(connections map[string][]string, start string, end string) []string {
	queue := []string{start}

	visited := make(map[string]bool)
	parent := make(map[string]string)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		visited[curr] = true

		if curr == end {
			path := []string{end}
			for parent[curr] != "" {
				curr = parent[curr]
				path = append([]string{curr}, path...)
			}
			return path
		}

		neighbors := connections[curr]
		for _, n := range neighbors {
			if !visited[n] {
				queue = append(queue, n)
				parent[n] = curr
			}
		}

	}
	return []string{start}
}

func ParseInput(input []string) map[string][]string {
	graph := map[string][]string{}

	for _, line := range input {
		data := strings.Split(line, ": ")
		connected := strings.Split(data[1], " ")
		graph[data[0]] = array.AppendIfNotPresent[string](graph[data[0]], connected...)
		for _, c := range connected {
			graph[c] = array.AppendIfNotPresent[string](graph[c], data[0])
		}
	}
	return graph
}
