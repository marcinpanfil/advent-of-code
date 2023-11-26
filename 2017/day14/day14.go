package aoc2017

import (
	day10 "advent-of-code/2017/day10"
	str_util "advent-of-code/util/strings"
	"fmt"
	"strconv"
	"strings"
)

func CalculateUsed(input string) int {
	hashes := BuildGrid(input)

	used := 0
	for _, hash := range hashes {
		used += strings.Count(hash, "1")
	}
	return used
}

func CalculateGroups(input string) int {
	hashes := BuildGrid(input)

	groupCount := 0
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			if string(hashes[j][i]) == "1" {
				// change the value of square j,i to "0"
				hashes[j] = str_util.ReplaceCharAt(hashes[j], i, "0")
				groupCount++

				queue := GetAdjecentUsed(hashes, i, j)
				for len(queue) > 0 {
					curX, curY := queue[0][0], queue[0][1]
					hashes[curY] = str_util.ReplaceCharAt(hashes[curY], curX, "0")
					queue = queue[1:]
					queue = append(queue, GetAdjecentUsed(hashes, curX, curY)...)
				}
			}
		}
	}
	return groupCount
}

func GetAdjecentUsed(hashes []string, curX int, curY int) [][2]int {
	// x, y
	directions := [][2]int{
		{0, 1},  // up
		{-1, 0}, // left
		{0, -1}, // down
		{1, 0},  // right
	}

	adj := [][2]int{}
	for _, dir := range directions {
		x := dir[0] + curX
		y := dir[1] + curY
		if x < 128 && x >= 0 && string(hashes[curY][x]) == "1" && x != curX {
			adj = append(adj, [2]int{x, y})
		}
		if y < 128 && y >= 0 && string(hashes[y][curX]) == "1" && y != curY {
			adj = append(adj, [2]int{x, y})
		}
	}
	return adj
}

func BuildGrid(input string) []string {
	hashes := []string{}
	for i := 0; i < 128; i++ {
		knotHash := day10.CalculateKnotHash(fmt.Sprintf("%s-%v", input, i))
		knotBin := ""
		for i := 0; i < len(knotHash); i++ {
			result, err := HexToBin(knotHash[i : i+1])
			if err != nil {
				panic(err)
			}
			knotBin += result
		}
		hashes = append(hashes, knotBin)
	}
	return hashes
}

func HexToBin(hex string) (string, error) {
	ui, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return fmt.Sprintf("%04b", ui), nil
}
