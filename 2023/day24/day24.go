package aoc2023

import (
	"advent-of-code/util/math"
	"strconv"
	"strings"
)

type Hailstone struct {
	px, py, pz float64
	vx, vy, vz float64
	slope      float64
}

var MIN = float64(200000000000000)
var MAX = float64(400000000000000)

func FindSingleThrow(input []string) int {
	hailstones := ParseInput(input)
	argXY := [][]float64{}
	argYZ := [][]float64{}

	for i := 0; i < 4; i++ {
		h1 := hailstones[i]
		h2 := hailstones[i+1]

		vy := h2.vy - h1.vy
		vx := h1.vx - h2.vx
		py := h1.py - h2.py
		px := h2.px - h1.px
		vz := h1.vz - h2.vz
		pz := h2.pz - h1.pz
		xy := h2.px*h2.vy - h2.py*h2.vx - h1.px*h1.vy + h1.py*h1.vx
		yz := h2.pz*h2.vy - h2.py*h2.vz - h1.pz*h1.vy + h1.py*h1.vz
		argXY = append(argXY, []float64{vy, vx, py, px, xy})
		argYZ = append(argYZ, []float64{vy, vz, py, pz, yz})
	}

	math.GaussianElimination(argXY)
	math.GaussianElimination(argYZ)

	lineLen := len(argXY[0]) - 1
	return int(argXY[0][lineLen]+0.5) + int(argXY[1][lineLen]+0.5) + int(argYZ[0][lineLen]+0.5)
}

func CalculateCollidePositions(input []string) int {
	hailstones := ParseInput(input)
	collideCount := 0
	for i := 0; i < len(hailstones)-1; i++ {
		h1 := hailstones[i]
		for j := i + 1; j < len(hailstones); j++ {
			h2 := hailstones[j]
			intersectionX := float64(0)
			intersectionY := float64(0)
			// parallel no intersection
			if h1.slope == h2.slope {
				continue
				// h1 is verticall
			} else if h1.slope == 0 {
				intersectionX = h1.px
				intersectionY = h2.slope*(intersectionX-h2.px) + h2.py
				// h2 is verticall
			} else if h2.slope == 0 {
				intersectionX = h2.px
				intersectionY = h1.slope*(intersectionX-h1.px) + h1.py
				// see README
			} else {
				intersectionX = (h1.py - h2.py - (h1.slope * h1.px) + (h2.slope * h2.px)) / (h2.slope - h1.slope)
				intersectionY = h1.slope*(intersectionX-h1.px) + h1.py
			}
			if HasIntersection(intersectionX, intersectionY, h1, h2) {
				collideCount++
			}
		}
	}

	return collideCount
}

func HasIntersection(iX, iY float64, h1, h2 Hailstone) bool {
	isInterPartOfX := (iX - h1.px) / h1.vx
	isInterPartOfY := (iX - h2.px) / h2.vx
	if isInterPartOfX > 0 && isInterPartOfY > 0 && iX >= MIN && iX <= MAX && iY >= MIN && iY <= MAX {
		return true
	}
	return false
}

func ParseInput(input []string) []Hailstone {
	hailstones := []Hailstone{}
	for _, line := range input {
		data := strings.Split(line, " @ ")
		posData := strings.Split(data[0], ", ")
		velData := strings.Split(data[1], ", ")
		px, _ := strconv.ParseFloat(posData[0], 64)
		py, _ := strconv.ParseFloat(posData[1], 64)
		pz, _ := strconv.ParseFloat(posData[2], 64)
		vx, _ := strconv.ParseFloat(velData[0], 64)
		vy, _ := strconv.ParseFloat(velData[1], 64)
		vz, _ := strconv.ParseFloat(velData[2], 64)
		if vx != 0 {
			hailstones = append(hailstones, Hailstone{px: px, py: py, pz: pz, vx: vx, vy: vy, vz: vz, slope: vy / vx})
		} else {
			hailstones = append(hailstones, Hailstone{px: px, py: py, pz: pz, vx: vx, vy: vy, vz: vz, slope: 0})
		}

	}
	return hailstones
}
