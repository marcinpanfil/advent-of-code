package aoc2023

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

type Mapping struct {
	destRangeStart   int
	sourceRangeStart int
	rangeLen         int
}

type SeedRange struct {
	start  int
	_range int
}

func FindMinLocation(input []string) int {
	seeds := GetSeeds(input[0])
	fmt.Println(seeds)
	seedToSoil := GetMap("seed-to-soil map", input)
	soilToFertilizer := GetMap("soil-to-fertilizer map", input)
	fertilizerToWater := GetMap("fertilizer-to-water map", input)
	waterToLight := GetMap("water-to-light map", input)
	lightToTemp := GetMap("light-to-temperature map", input)
	tempToHumidity := GetMap("temperature-to-humidity map", input)
	humidityToLoc := GetMap("humidity-to-location map", input)

	min := math.MaxInt64
	for _, seed := range seeds {
		soil := GetFromMap(seedToSoil, seed)
		fertilizer := GetFromMap(soilToFertilizer, soil)
		water := GetFromMap(fertilizerToWater, fertilizer)
		light := GetFromMap(waterToLight, water)
		temp := GetFromMap(lightToTemp, light)
		humidity := GetFromMap(tempToHumidity, temp)
		loc := GetFromMap(humidityToLoc, humidity)
		if loc < min {
			min = loc
		}
	}
	return min
}

func FindMinLocationForRanges(input []string) int {
	seedRanges := GetSeedRanges(input[0])
	seedToSoil := GetMap("seed-to-soil map", input)
	soilToFertilizer := GetMap("soil-to-fertilizer map", input)
	fertilizerToWater := GetMap("fertilizer-to-water map", input)
	waterToLight := GetMap("water-to-light map", input)
	lightToTemp := GetMap("light-to-temperature map", input)
	tempToHumidity := GetMap("temperature-to-humidity map", input)
	humidityToLoc := GetMap("humidity-to-location map", input)

	var waitGr sync.WaitGroup
	var mutex sync.Mutex

	min := math.MaxInt64
	for _, seedRange := range seedRanges {
		waitGr.Add(1)
		start := seedRange.start
		stop := seedRange.start + seedRange._range
		go func() {
			rangeMin := math.MaxInt64
			for seed := start; seed < stop; seed++ {
				soil := GetFromMap(seedToSoil, seed)
				fertilizer := GetFromMap(soilToFertilizer, soil)
				water := GetFromMap(fertilizerToWater, fertilizer)
				light := GetFromMap(waterToLight, water)
				temp := GetFromMap(lightToTemp, light)
				humidity := GetFromMap(tempToHumidity, temp)
				loc := GetFromMap(humidityToLoc, humidity)
				if loc < rangeMin {
					rangeMin = loc
				}
			}
			mutex.Lock()
			if rangeMin < min {
				min = rangeMin
			}
			mutex.Unlock()
			waitGr.Done()
		}()
	}
	waitGr.Wait()

	return min
}

func GetSeedRanges(seedLine string) []SeedRange {
	seedRanges := []SeedRange{}
	seedsAsStr := strings.Split(seedLine[7:], " ")
	for i := 0; i < len(seedsAsStr); i += 2 {
		start, _ := strconv.Atoi(seedsAsStr[i])
		_range, _ := strconv.Atoi(seedsAsStr[i+1])
		seedRanges = append(seedRanges, SeedRange{start: start, _range: _range})
	}
	return seedRanges
}

func GetSeeds(seedLine string) []int {
	seeds := []int{}
	seedsAsStr := strings.Split(seedLine[7:], " ")
	for _, s := range seedsAsStr {
		seed, _ := strconv.Atoi(s)
		seeds = append(seeds, seed)
	}
	return seeds
}

func GetMap(mapName string, input []string) []Mapping {
	isMapping := false
	mapping := []Mapping{}
	for _, line := range input {
		if strings.HasPrefix(line, mapName) {
			isMapping = true
		} else if isMapping && len(line) == 0 {
			return mapping
		} else if isMapping {
			dataAsStr := strings.Split(line, " ")
			destRangeStart, _ := strconv.Atoi(dataAsStr[0])
			sourceRangeStart, _ := strconv.Atoi(dataAsStr[1])
			rangeLen, _ := strconv.Atoi(dataAsStr[2])
			mapping = append(mapping, Mapping{destRangeStart: destRangeStart, sourceRangeStart: sourceRangeStart, rangeLen: rangeLen})
		}
	}
	return mapping
}

func GetFromMap(mapping []Mapping, value int) int {
	for _, entry := range mapping {
		if entry.sourceRangeStart+entry.rangeLen-1 >= value && entry.sourceRangeStart <= value {
			return entry.destRangeStart + (value - entry.sourceRangeStart)
		}
	}
	return value
}
