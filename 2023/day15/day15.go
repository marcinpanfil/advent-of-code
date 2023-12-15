package aoc2023

import (
	"strconv"
	"strings"
)

func CalculateHashValue(input []string) int {
	result := 0
	for _, step := range input {
		result += CalculateHash(step)
	}
	return result
}

func CalculateHash(step string) int {
	curr := 0
	for _, c := range step {
		value := int(c)
		curr += value
		curr *= 17
		curr %= 256
	}
	return curr
}

type Box struct {
	lenses []Label
}

type Label struct {
	name  string
	focal int
}

func (b *Box) addLense(label Label) {
	for i := range b.lenses {
		l := b.lenses[i]
		if l.name == label.name {
			l.focal = label.focal
			b.lenses[i] = l
			return
		}
	}
	b.lenses = append(b.lenses, label)

}

func (b *Box) removeLense(name string) {
	for i, l := range b.lenses {
		if l.name == name {
			b.lenses = append(b.lenses[:i], b.lenses[i+1:]...)
			return
		}
	}
}

func FillLensSlots(input []string) int {
	boxes := make([]Box, 256)
	for _, step := range input {
		if strings.Contains(step, "-") {
			data := strings.Split(step, "-")
			curr := CalculateHash(data[0])
			boxes[curr].removeLense(data[0])
		} else if strings.Contains(step, "=") {
			data := strings.Split(step, "=")
			focal, _ := strconv.Atoi(data[1])
			curr := CalculateHash(data[0])
			label := Label{name: data[0], focal: focal}
			boxes[curr].addLense(label)
		}
	}
	result := 0

	for i, box := range boxes {
		for j, lense := range box.lenses {
			result += (i + 1) * (j + 1) * lense.focal
		}
	}

	return result
}
