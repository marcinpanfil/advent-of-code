package aoc2023

import (
	maps "advent-of-code/util/map"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
	z int
}

func newPosition(data string) Position {
	cordinates := strings.Split(data, ",")
	x, ex := strconv.Atoi(cordinates[0])
	y, ey := strconv.Atoi(cordinates[1])
	z, ez := strconv.Atoi(cordinates[2])
	if ex != nil || ey != nil || ez != nil {
		fmt.Println("Error")
	}
	return Position{x: x, y: y, z: z}
}

type Brick struct {
	id        int
	positions []Position
	minLvl    int
	maxLvl    int
}

func newBrick(id int, start, end Position) Brick {
	pos := []Position{}
	for x := start.x; x <= end.x; x++ {
		for y := start.y; y <= end.y; y++ {
			for z := start.z; z <= end.z; z++ {
				pos = append(pos, Position{x: x, y: y, z: z})
			}
		}
	}
	if end.z < start.z {
		fmt.Println("Error")
	}
	return Brick{id: id, positions: pos, minLvl: start.z, maxLvl: end.z}
}

func (b *Brick) contains(pos Position) bool {
	for _, p := range b.positions {
		if p.x == pos.x && p.y == pos.y && p.z == pos.z {
			return true
		}
	}
	return false
}

func (b *Brick) fallDown(diff int) {
	for i := range b.positions {
		p := b.positions[i]
		p.z -= diff
		b.positions[i] = p
	}
	b.minLvl -= diff
	b.maxLvl -= diff
}

func SolutionPart1(input []string) int {
	data := ParseInput(input)
	afterFall := FallAllDown(data)
	return CountDisintegrated(afterFall)
}

func SolutionPart2(input []string) int {
	data := ParseInput(input)
	return FindSumOfFallenBricks(data)
}

func ParseInput(input []string) []Brick {
	bricks := []Brick{}
	for i, line := range input {
		data := strings.Split(line, "~")
		start := newPosition(data[0])
		end := newPosition(data[1])
		brick := newBrick(i, start, end)
		bricks = append(bricks, brick)
	}
	sort.Slice(bricks, func(i, j int) bool {
		brick1 := bricks[i]
		brick2 := bricks[j]
		return brick1.minLvl < brick2.minLvl
	})
	return bricks
}

func FallAllDown(bricks []Brick) []Brick {
	// mapping of x, y cordinates to z cordinate
	xyMappingToZ := map[[2]int]int{}
	for i := 0; i < len(bricks); i++ {
		brick := bricks[i]
		minPossibleZ := FindMinimalZ(brick, xyMappingToZ)
		diffZ := brick.minLvl - minPossibleZ
		if diffZ > 0 {
			brick.fallDown(diffZ)
			bricks[i] = brick
		}
		for _, position := range brick.positions {
			if position.z == brick.minLvl {
				xy := [2]int{position.x, position.y}
				xyMappingToZ[xy] = minPossibleZ + (brick.maxLvl - brick.minLvl)
			}
		}
	}
	return bricks
}

func FindMinimalZ(brick Brick, xyMappingToZ map[[2]int]int) int {
	minPossibleZ := -1
	for _, position := range brick.positions {
		if position.z == brick.minLvl {
			xy := [2]int{position.x, position.y}
			zPos := xyMappingToZ[xy]
			if zPos == 0 && minPossibleZ < 0 {
				// no block found in x y position, so z is minimal (1)
				minPossibleZ = 1
			} else if zPos+1 == brick.minLvl {
				// the found x y position is smaller than brick min z,
				// so it's not possible to move the brick
				minPossibleZ = brick.minLvl
				break
			} else if zPos >= minPossibleZ {
				// the found x y position is bigger than the current min z,
				// so this is the new min Z
				minPossibleZ = zPos + 1
			}
		}
	}
	return minPossibleZ
}

func CountDisintegrated(bricks []Brick) int {
	supported, supports := BuildSupportMatrix(bricks)

	candidateToDisintegrate := map[int]bool{}
	onlySupporters := map[int]bool{}

	for _, sup := range supported {
		if len(sup) > 1 {
			for k := range sup {
				candidateToDisintegrate[k] = true
			}
		} else if len(sup) == 1 {
			// can't disintegrate if brick is the only supporter
			for k := range sup {
				onlySupporters[k] = true
			}
		}
	}
	for support, supported := range supports {
		if len(supported) == 0 {
			candidateToDisintegrate[support] = true
		}
	}
	for supporter := range onlySupporters {
		delete(candidateToDisintegrate, supporter)
	}

	return len(candidateToDisintegrate)
}

func FindSumOfFallenBricks(bricks []Brick) int {
	bricks = FallAllDown(bricks)
	supported, supports := BuildSupportMatrix(bricks)

	counter := 0
	for _, brick := range bricks {
		cpSupported := make(map[int]map[int]bool, len(supported))
		maps.Copy(supported, &cpSupported)

		destroyed := map[int]bool{}
		queue := []int{brick.id}
		for len(queue) > 0 {
			current := queue[0]
			destroyed[current] = true
			queue = queue[1:]
			currSupports := supports[current]
			for k := range currSupports {
				supportedBy := cpSupported[k]
				delete(supportedBy, current)
				cpSupported[k] = supportedBy
				if len(cpSupported[k]) == 0 && !destroyed[k] {
					queue = append(queue, k)
				}
			}
		}

		for _, sup := range cpSupported {
			if len(sup) == 0 {
				counter++
			}
		}
	}
	return counter
}

func BuildSupportMatrix(bricks []Brick) (map[int]map[int]bool, map[int]map[int]bool) {
	supported := map[int]map[int]bool{}
	supports := map[int]map[int]bool{}
	for i := 0; i < len(bricks); i++ {
		brick1 := bricks[i]
		supports[brick1.id] = map[int]bool{}
		for j := i; j < len(bricks); j++ {
			brick2 := bricks[j]
			if brick1.maxLvl+1 == brick2.minLvl {
				for _, p := range brick1.positions {
					if brick2.contains(Position{x: p.x, y: p.y, z: p.z + 1}) {
						UpdateSupportLinkage(brick1.id, brick2.id, supported, supports)
						break
					}
				}
			}
		}
	}
	return supported, supports
}

func UpdateSupportLinkage(b1Id, b2Id int, supported map[int]map[int]bool, supports map[int]map[int]bool) {
	sup := supported[b2Id]
	if sup == nil {
		supported[b2Id] = map[int]bool{b1Id: true}
	} else {
		supported[b2Id][b1Id] = true
	}
	supports[b1Id][b2Id] = true
}
